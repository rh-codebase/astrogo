/*
 *   Wraps the NOVAS library and any other ephemeris related functions into
 *   a simple class.
 *   For most cases one instantiates this class with a source, and
 *   alternatively a fixed station
 *           e := NewEphemeris("3c273","BIMA-9")
 *   after which the time should be set (time needs to be set *after* a source)
 *           e.setMJD()
 *   and the RA/DEC or AZ/EL can be retrieved
 *           double ra=e.getRa()
 *           double az=e.getAz()
 *   Caveat: for more accurate observations you need to feed the ephemeris
 *   a more detailed atmosphere description (pressure, temperature, humidity)
 *   as well as an observing frequency.
 *
 *   Note that solar system calculations are often done in TT, not JD,
 *   where TT = JD + deltat
 *
 *   Known issues:
 *   - solar sytem objects are done in TT, not UTC
 *     if you then take known RA/DEC and stuff it into a catalog,
 *     you will be off by deltat (about 65 sec) proper motion
 *
 *
 */
package ephemeris

import (
	"errors"
	"fmt"
	"strings"
	"time"

	au "github.com/rh-codebase/astrogo/astrounit"
	nov "github.com/rh-codebase/novasgo/novas"
	"gopkg.in/yaml.v2"
)

type Wx struct {
	AtmTemperature au.Temperature
	AtmPressure    au.Pressure
	RelHumidityPct float64
}

type StarInfo struct {
	Name            string
	Catalog         string
	StarNum         int64
	Ra_hr           float64
	Dec_deg         float64
	PMRA_masPerYr   float64
	PMDEC_masPerYr  float64
	Parallax_mas    float64
	RadVel_kmPerSec float64
}

type Location struct {
	Latitude  au.Angle
	Longitude au.Angle
	Height    au.Length
}

type RaDec struct {
	Ra_hr   float64 `yaml:"ra"`
	Dec_deg float64 `yaml:"dec"`
}

const (
	fk5 = "FK5"
	// solar system body strings
	Sun     = "sun"
	Moon    = "moon"
	Mercury = "mercury"
	Venus   = "venus"
	Earth   = "earth"
	Mars    = "mars"
	Jupiter = "jupiter"
	Saturn  = "saturn"
	Neptune = "neptune"
	Uranus  = "uranus"
	Pluto   = "pluto"
)

var (
	wx                     Wx
	site                   nov.OnSurface
	radec                  RaDec
	location               Location
	relhum                 float64 // percent
	recompute              bool
	doRefract              bool
	is_cat                 bool
	use_source             bool
	observefreq            float64 // Hz
	tjd                    float64
	mjd                    float64
	deltat, xxpole, y_pole float64
	is_eph                 bool
	dra, ddec, daz, del    float64
	doppler                float64
	offsetMode             int16
	sourceName             string
	ra2000                 float64
	fixed                  bool
)

// SetWeather sets the Wx structure. rh is relative humidity in percent
func SetWeather(ap au.Pressure, at au.Temperature, rh float64) {
	recompute = true

	wx.AtmPressure = ap
	wx.AtmTemperature = at
	wx.RelHumidityPct = rh
	site.Pressure = ap.ToMillibar().Value
	site.Temperature = at.ToCelsius().Value
	// percent
	// @todo warn here if input values were bad
	relhum = rh
}

// SetFreq sets the observation frequency in Hz. That is, the frequency of the
// radiation being collected by a sensor.
func SetFreq(freq float64) {
	recompute = true
	observefreq = freq // Hz
}

// GetFreq returns the frequency of observed radiation in Hz.
func GetFreq() float64 {
	return observefreq // Hz
}

// SetRefraction controls whether or not to apply refraction to the
// elevation obtained from NOVAS.
func SetRefraction(refract bool) {
	recompute = true
	doRefract = refract
}

// Refract computes and applies the refraction angle given the elevation angle
// computed by NOVAS. Functions: SetWeather(), SetFreq(), SetLocation()
// and SetRefraction() must be called beforehand
func Refract(el au.Angle) (au.Angle, error) {
	if doRefract {
		ra, err := au.ComputeRefractionCorrection(wx.AtmTemperature, wx.AtmPressure,
			wx.RelHumidityPct, el, GetFreq(), location.Height)
		if err != nil {
			return el, err
		}
		corrected := el.Add(ra)
		return corrected, nil
	}
	return el, nil
}

func getSourceFromCatalog(src string, bsc *BSC) (StarInfo, error) {
	var starInfo StarInfo
	star, err := bsc.GetSource(src)
	if err != nil {
		return starInfo, err
	}
	starInfo.Name = src
	starInfo.Catalog = "BSC"
	starInfo.StarNum = 1
	starInfo.Ra_hr = star.RA.Hour().Value
	starInfo.Dec_deg = star.DEC.Degree().Value
	starInfo.PMRA_masPerYr = star.PMRA.MilliArcSecond().Value
	starInfo.PMDEC_masPerYr = star.PMDEC.MilliArcSecond().Value
	starInfo.Parallax_mas = 0.0
	starInfo.RadVel_kmPerSec = 0.0

	return starInfo, nil
}

func isPlanet(name string) bool {
	switch strings.ToLower(name) {
	case Sun, Moon, Mercury, Venus, Mars, Jupiter, Neptune, Uranus, Pluto:
		return true
	default:
		return false
	}
}

// isRaDec is a helper to determine if the string value came from
// a serilazied radec structure
func isRaDec(rd string) bool {
	err := yaml.UnmarshalStrict([]byte(rd), &radec)
	if radec.Ra_hr < 0.0 || radec.Ra_hr > 24.0 {
		fmt.Printf("Invalid Ra value. Must be in hours: %f", radec.Ra_hr)
		return false
	}
	if err != nil {
		return false
	} else {
		return true
	}
}

// SImpleTrack returns a function to allow updating a source's position in
// az.el coordiantes based on time. OnSurface represents the observer's location
// on Earth and the sourcename must be in the BSC catalog.
func SimpleTrack(si nov.OnSurface, sourceName string, bsc *BSC) (func(time.Time) (float64, float64, error), error) {

	src := strings.ToLower(sourceName)
	// not really number of leapseconds. this should just be called 'leap'
	// since it represents the number of seconds TAI is ahead of UTC.
	var Leap int16 = 37
	var ut1Utc float64 = -0.17442

	var catEntry nov.CatEntry
	var source nov.Object
	planet := isPlanet(src)
	if planet {
		switch strings.ToLower(src) {
		case Sun:
			nov.MakeObject(0, 10, Sun, &catEntry, &source)
		case Moon:
			nov.MakeObject(0, 11, Moon, &catEntry, &source)
		case Mercury:
			nov.MakeObject(0, 1, Mercury, &catEntry, &source)
		case Venus:
			nov.MakeObject(0, 2, Venus, &catEntry, &source)
		case Mars:
			nov.MakeObject(0, 4, Mars, &catEntry, &source)
		case Jupiter:
			nov.MakeObject(0, 5, Jupiter, &catEntry, &source)
		case Saturn:
			nov.MakeObject(0, 6, Saturn, &catEntry, &source)
		case Uranus:
			nov.MakeObject(0, 7, Uranus, &catEntry, &source)
		case Neptune:
			nov.MakeObject(0, 8, Neptune, &catEntry, &source)
		case Pluto:
			nov.MakeObject(0, 9, Pluto, &catEntry, &source)
		default:
			emsg := fmt.Sprintf("Unknown Planet name: %s", sourceName)
			return nil, errors.New(emsg)
		}
	} else if isRaDec(src) {
		fmt.Println("SimpleTrack: Making RaDec source. Ra_hr= ", radec.Ra_hr, " Dec_deg= ", radec.Dec_deg)
		// NOTE: 2nd field must be "BSC".
		nov.MakeCatEntry("RaDec", "BSC", 1, radec.Ra_hr, radec.Dec_deg,
			0.0, 0.0, 0.0, 0.0, &catEntry)
	} else { // last gasp to see if src is in a catalog
		//var catEntry nov.CatEntry
		// J2000 AlpBoo    14:15:39.70 +19:10:57.0 -0:0:00.0729 -0:0:01.998 # -0.0
		starInfo, err := getSourceFromCatalog(sourceName, bsc)
		if err != nil {
			return nil, err
		}
		nov.MakeCatEntry(starInfo.Name, starInfo.Catalog, starInfo.StarNum,
			starInfo.Ra_hr, starInfo.Dec_deg,
			starInfo.PMRA_masPerYr, starInfo.PMDEC_masPerYr,
			starInfo.Parallax_mas, starInfo.RadVel_kmPerSec, &catEntry)
	}

	return func(ti time.Time) (az float64, el float64, err error) {
		year := int16(ti.Year())
		month := int16(ti.Month())
		day := int16(ti.Day())
		hr := ti.Hour()
		min := ti.Minute()
		sec := ti.Second()
		ns := ti.Nanosecond()
		hour := float64(hr) + float64(min)/60. + (float64(sec)+float64(ns)/1e9)/3600.
		jdUTC := nov.JulianDate(year, month, day, hour)
		//		fmt.Println(jdUTC)
		jdTT := jdUTC + (float64(Leap)+32.184)/86400.
		jdUT1 := jdUTC + ut1Utc/86400.0
		deltaT := 32.184 + float64(Leap) - ut1Utc
		var ra, dec, dis float64
		var accuracy int16 = 0
		if planet {
			err = nov.TopoPlanet(jdTT, &source, deltaT, &si, accuracy, &ra, &dec, &dis)
			if err != nil {
				return az, el, err
			}
		} else {
			fmt.Println("SimpleTrack: catentry: ", catEntry)
			err = nov.TopoStar(jdTT, deltaT, &catEntry, &si, accuracy, &ra, &dec)
			if err != nil {
				return az, el, err
			}
		}
		var zd, rar, decr float64
		doRefraction := int16(0)
		//fmt.Println("jdUT1, deltaT, ra, dec: ", jdUT1, deltaT, ra, dec)
		nov.Equ2hor(jdUT1, deltaT, accuracy, 0.0, 0.0, &si, ra, dec, doRefraction, &zd, &az, &rar, &decr)

		return az, 90.0 - zd, nil
	}, nil

}

/*
func (l *Location) GetLatitude() au.Angle {
	return l.Latitude
}

func (l *Location) GetLongitude() au.Angle {
	return l.Longitude
}
*/

/*
func (l *Location) GetHeight() au.Length {
	return l.Height
  }
*/

// TODO: Implement or remove
func SetSourceCat(sourceName, catalogName string) {

}

// SetLocation sets a Location structure.
func SetLocation(loc Location) {
	recompute = true
	site.Latitude = loc.Latitude.Degree().Value
	site.Longitude = loc.Longitude.Degree().Value
	site.Height = loc.Height.Meter().Value
	location = loc
}

// GetLocations returns the location structure
func GetLocation() Location {
	return location
}
