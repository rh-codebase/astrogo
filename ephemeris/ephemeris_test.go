package ephemeris

import (
	"fmt"
	"testing"
	"time"

	au "github.com/rh-codebase/astrogo/astrounit"
	nov "github.com/rh-codebase/novasgo/novas"
)

/*
func TestGetLatitude(t *testing.T) {
	loc = Location{1.0, 2.0, 3.0}
	lat := loc.GetLatittude()

	fmt.Println("lat= ", lat)
}
*/

func TestSimpleTrack(t *testing.T) {
	//var a au.Angle
	//fmt.Println("a: ", a.Degree().Value)

	bsc := make(BSC)
	cs := CatalogSource{"BSC", "brightSourceCatalog.yml"}

	AddCatalog(cs)
	bsc.LoadCatalogs()
	fmt.Println("bsc[AlpBoo]: ", bsc["alpboo"])

	var si nov.OnSurface
	nov.MakeOnSurface(37.2339, -118.282, 1222., 0.0, 0.0, &si)
	sourceName := "Sun"
	sunAzEl, err := SimpleTrack(si, sourceName, &bsc)
	if err != nil {
		fmt.Println("Error SimpleTrack for Sun")
		t.Fail()
	} else {
		for idx := 0; idx < 5; idx++ {
			ti := time.Now().UTC()
			az, el, err := sunAzEl(ti)
			if err != nil {
				fmt.Println("NextAzEl error: ", err)
				t.Fail()
			}
			fmt.Printf("[%v]%s: az, el= %6.3f, %5.3f\n", ti, sourceName, az, el)
			time.Sleep(1 * time.Second)
		}
	}

	jupAzEl, err := SimpleTrack(si, "Jupiter", &bsc)
	if err != nil {
		fmt.Println("Error SimpleTrack for Jupiter")
		t.Fail()
	} else {
		for idx := 0; idx < 5; idx++ {
			ti := time.Now().UTC()
			az, el, err := jupAzEl(ti)
			if err != nil {
				fmt.Println("NextAzEl error: ", err)
				t.Fail()
			}
			fmt.Printf("[%v]%s: az, el= %6.3f, %5.3f\n", ti, "Jupiter", az, el)
			time.Sleep(1 * time.Second)
		}
	}

	moonAzEl, err := SimpleTrack(si, "Moon", &bsc)
	if err != nil {
		fmt.Println("Error SimpleTrack for Moon")
		t.Fail()
	} else {
		for idx := 0; idx < 5; idx++ {
			ti := time.Now().UTC()
			az, el, err := moonAzEl(ti)
			if err != nil {
				fmt.Println("NextAzEl error: ", err)
				t.Fail()
			}
			fmt.Printf("[%v]%s: az, el= %6.3f, %5.3f\n", ti, "Moon", az, el)
			time.Sleep(1 * time.Second)
		}
	}

	starAzEl, err := SimpleTrack(si, "alpboo", &bsc)
	if err != nil {
		fmt.Println("Error SimpleTrack for AlpBoo: ", err)
		t.Fail()
	} else {
		for idx := 0; idx < 5; idx++ {
			ti := time.Now().UTC()
			az, el, err := starAzEl(ti)
			if err != nil {
				fmt.Println("NextAzEl error: ", err)
				t.Fail()
			}
			fmt.Printf("[%v]%s: az, el= %6.3f, %5.3f\n", ti, "AlpBoo", az, el)
			time.Sleep(1 * time.Second)
		}
	}

	// Now a Star
}

func TestRefract(t *testing.T) {
	airT, _ := au.NewTemperature(au.Celsius, 10.0)
	airP := au.NewPressure(au.Millibar, 1013.0)
	rh := 50.0
	freq := 1.e6

	SetWeather(airP, airT, rh)
	SetFreq(freq)
	SetRefraction(true)
	var loc Location
	loc.Latitude = au.NewAngle(au.Degree, 37.0)
	loc.Longitude = au.NewAngle(au.Degree, -118.0)
	loc.Height = au.NewLength(au.Meter, 3.0)
	SetLocation(loc)
	el := au.NewAngle(au.Degree, 45.0)
	cel, err := Refract(el)
	if err != nil {
		fmt.Println("Refract returned error: ", err)
		t.Fail()
	}
	fmt.Printf("before El: %8.5f, After: %8.5f\n", el.Degree().Value, cel.Degree().Value)
}

func TestCatalog(t *testing.T) {
	bsc := make(BSC)
	cs := CatalogSource{"BSC", "brightSourceCatalog.yml"}

	AddCatalog(cs)
	bsc.LoadCatalogs()
	fmt.Println("bsc[AlpBoo]: ", bsc["alpboo"])
	ra := bsc["alpboo"].RA
	fmt.Println("bsc[AlpBoo] ra: ", ra.Hour().Value, ra.SexagesimalHMS())

}
