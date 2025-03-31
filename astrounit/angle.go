// Angle type
package astrounit

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	at "github.com/rh-codebase/astrogo/astrotime"
)

type AngleUnit int

const (
	// Enums to identify Angle type
	_ AngleUnit = iota
	Radian
	Degree
	MilliRadian
	ArcMinute
	ArcSecond
	MilliArcSecond
	Hour

	// Angle unit strings
	RadianStr         = "rad"
	DegreeStr         = "deg"
	MilliRadianStr    = "mrad"
	ArcMinuteStr      = "arcmin"
	ArcSecondStr      = "arcsec"
	MilliArcSecondStr = "mas"
	HourStr           = "hr"
	DMSstr            = "%.2d:%.2d:%07.4f"
	DMSnegStr         = "-%.2d:%.2d:%07.4f"
	HMSstr            = "%.2d:%.2d:%07.4f"
	HMSnegStr         = "-%.2d:%.2d:%07.4f"
)

// Angle supports a value and associated unit
type Angle struct {
	Unit  AngleUnit `yaml:"unit"`
	Value float64   `yaml:"value"`
}

type AngleVector []Angle

// HMS represents HH:MM:SS.ss
type HMS struct {
	Hr  float64 `yaml:"hr"`
	Min float64 `yaml:"min"`
	Sec float64 `yaml:"sec"`
}

// HMS represents DD:MM:SS.ss
type DMS struct {
	Deg float64 `yaml:"hr"`
	Min float64 `yaml:"min"`
	Sec float64 `yaml:"sec"`
}

// pasrseSexagesimal returns the three tokens from a string in the form
// (sign)xx:yy:ss.ss where (sign) is optional for positive values or
// '+' or '-'. Ex: -65:34:37.9 and +0:0:00.0032. The first would return
// -65.0, 34.0, 37.9, nil
func parseSexagesimal(sex string) (float64, float64, float64, error) {
	tokens := strings.Split(sex, ":")
	if len(tokens) != 3 {
		emsg := fmt.Sprintf("Invalid Sexagesimal input string: %s", sex)
		return 0.0, 0.0, 0.0, errors.New(emsg)
	}

	sign := string(tokens[0][0])
	var fac float64
	var hrstr string
	if sign == "+" {
		fac = 1.0
		hrstr = string(tokens[0][1:])
	} else if sign == "-" {
		fac = -1.0
		hrstr = string(tokens[0][1:])
	} else {
		fac = 1.0
		hrstr = string(tokens[0])
	}
	//fmt.Println("fac: ", fac, " hrstr: ", hrstr)

	hr, err := strconv.ParseFloat(hrstr, 64)
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}
	min, err := strconv.ParseFloat(tokens[1], 64)
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}
	sec, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}

	return fac * hr, fac * min, fac * sec, nil
}

// NewHMS returns an HMS angle given a sexgisamal string of the form:
// (sign)hh:mm:ss.ss where (sign) is optional for positive angles or
// '+' or '-'. Ex: -65:34:37.9 and +0:0:00.0032
func NewHMS(hms string) (HMS, error) {
	var res HMS
	hr, min, sec, err := parseSexagesimal(hms)
	if err != nil {
		return res, err
	}
	res.Hr = hr
	res.Min = min
	res.Sec = sec
	return res, nil
}

// NewDMS returns an DMS angle given a sexgisamal string of the form:
// (sign)hh:mm:ss.ss where (sign) is optional for positive angles or
// '+' or '-'
func NewDMS(dms string) (DMS, error) {
	var res DMS
	deg, min, sec, err := parseSexagesimal(dms)
	if err != nil {
		return res, err
	}
	res.Deg = deg
	res.Min = min
	res.Sec = sec
	return res, nil
}

// NewAngle returns an Angle in the specified units and value
func NewAngle(au AngleUnit, v float64) Angle {
	var a Angle
	a.Unit = au
	a.Value = v
	return a
}

// NewAngleDMS returns an Angle given sexagesimal
func NewAngleDMS(s string) (Angle, error) {
	var a Angle
	dms, err := NewDMS(s)
	if err != nil {
		return a, err
	}
	return dms.Angle(), nil
}

// NewAngleHMS returns an Angle given sexagesimal
func NewAngleHMS(s string) (Angle, error) {
	var a Angle
	hms, err := NewHMS(s)
	if err != nil {
		return a, err
	}
	return hms.Angle(), nil
}

// UnitString returns the units as a string for the given Angle
func (a Angle) UnitString() string {
	var as string
	switch a.Unit {
	case Radian:
		as = RadianStr
	case Degree:
		as = DegreeStr
	case MilliRadian:
		as = MilliRadianStr
	case ArcMinute:
		as = ArcMinuteStr
	case ArcSecond:
		as = ArcSecondStr
	case MilliArcSecond:
		as = MilliArcSecondStr
	case Hour:
		as = HourStr
	}
	return as
}

// SexagesimalDMS returns the sexagesimal notation for the angle in DMS
func (a *Angle) SexagesimalDMS() string {
	dms := a.DMS()
	return dms.UnitString()
}

// SexagesimalHMS returns the sexagesimal notation for the angle in HMS
func (a *Angle) SexagesimalHMS() string {
	hms := a.HMS()
	return hms.UnitString()
}

// UnitString returns the units for a DMS type
func (dms *DMS) UnitString() string {
	if dms.Deg < 0.0 || dms.Min < 0.0 || dms.Sec < 0.0 {
		return fmt.Sprintf(DMSnegStr, int(math.Abs(dms.Deg)),
			int(math.Abs(dms.Min)), math.Abs(dms.Sec))
	} else {
		return fmt.Sprintf(DMSstr, int(dms.Deg), int(dms.Min), dms.Sec)
	}
}

// UnitString returns the units for an HMS type
func (hms *HMS) UnitString() string {
	if hms.Hr < 0.0 || hms.Min < 0.0 || hms.Sec < 0.0 {
		return fmt.Sprintf(HMSnegStr, int(math.Abs(hms.Hr)),
			int(math.Abs(hms.Min)), math.Abs(hms.Sec))
	} else {
		return fmt.Sprintf(HMSstr, int(hms.Hr), int(hms.Min), hms.Sec)
	}
}

// HMS converts Angle to a HMS type
// No checks for NaN or +/-Inf
func (a Angle) HMS() HMS {
	var hms HMS
	ah := a.Hour()
	hms.Hr = float64(int(ah.Value))
	hms.Min = float64(int((ah.Value - hms.Hr) * at.MinutePerHour))
	hms.Sec = (ah.Value-hms.Hr)*at.SecondPerHour - hms.Min*at.SecondPerMinute
	return hms
}

// DMS converts Angle to a DMS type
// No checks for NaN or +/-Inf
func (a Angle) DMS() DMS {
	var dms DMS
	ad := a.Degree()
	dms.Deg = float64(int(ad.Value))
	dms.Min = float64(int((ad.Value - dms.Deg) * MinutePerDegree))
	dms.Sec = (ad.Value-dms.Deg)*SecondPerDegree - dms.Min*at.SecondPerMinute
	return dms
}

// Angle converts DMS to an Angle.
// If any element of DMS < 0.0, then the quantity is
// taken as < 0.0. That is, negative degrees will not be added to positive
// minutes. Ex. DMS{-1,30,0} => -1.5deg not -.5deg
// No checks for NaN or +/-Inf
func (dms DMS) Angle() Angle {
	var a Angle
	a.Unit = Radian
	if dms.Deg < 0.0 || dms.Min < 0.0 || dms.Sec < 0.0 {
		a.Value = -(math.Abs(dms.Deg)*RadianPerDegree + math.Abs(dms.Min)*RadianPerMinute +
			math.Abs(dms.Sec)*RadianPerSecond)
	} else {
		a.Value = dms.Deg*RadianPerDegree + dms.Min*RadianPerMinute +
			dms.Sec*RadianPerSecond
	}
	return a
}

// HMS converts DMS to HMS.  If any element of DMS < 0.0, then
// the quantity is taken as < 0.0. That is, negative degrees will
// not be added to positive minutes. Ex. DMS{-1,30,0} => -1.5deg not -.5deg
// No checks for NaN or +/-Inf
func (dms DMS) HMS() HMS {
	return dms.Angle().HMS()
}

// DMS converts HMS to DMS.  If any element of DMS < 0.0, then
// the quantity is taken as < 0.0. That is, negative degrees will
// not be added to positive minutes. Ex. DMS{-1,30,0} => -1.5deg not -.5deg
// No checks for NaN or +/-Inf.
// No side-effects on HMS
func (hms HMS) DMS() DMS {
	return hms.Angle().DMS()
}

// Radian onverts HMS to an Angle with Radian units.
// If any element of HMS < 0.0, then the quantity is
// taken as < 0.0. That is, negative hours will not be added to positive
// minutes. Ex. HMS{-1,30,0} => -1.5hr not -.5hr
// No checks for NaN or +/-Inf
func (hms HMS) Angle() Angle {
	var a Angle
	a.Unit = Radian
	if hms.Hr < 0.0 || hms.Min < 0.0 || hms.Sec < 0.0 {
		decHrs := math.Abs(hms.Hr) + math.Abs(hms.Min)/at.MinutePerHour +
			math.Abs(hms.Sec)/at.SecondPerHour
		a.Value = -decHrs * DegreePerHour * RadianPerDegree
	} else {
		decHrs := hms.Hr + hms.Min/at.MinutePerHour +
			hms.Sec/at.SecondPerHour
		a.Value = decHrs * DegreePerHour * RadianPerDegree
	}
	return a
}

// Radian return an Angle in Radian units.
// No checks for NaN or +/-Inf
func (a Angle) Radian() Angle {
	var aa Angle
	aa.Unit = Radian
	switch a.Unit {
	case Radian:
		aa.Value = a.Value
	case MilliRadian:
		aa.Value = a.Value / MilliRadianPerRadian
	case Degree:
		aa.Value = a.Value * RadianPerDegree
	case ArcMinute:
		aa.Value = a.Value / ArcMinutePerDegree * RadianPerDegree
	case ArcSecond:
		aa.Value = a.Value / ArcSecondPerDegree * RadianPerDegree
	case MilliArcSecond:
		aa.Value = a.Value / ArcSecondPerDegree * RadianPerDegree / 1000.
	case Hour:
		aa.Value = a.Value * DegreePerHour * RadianPerDegree
	}
	return aa
}

// MilliRadian returns an Angle in MilliRadians units.
// No checks for NaN or +/-Inf
func (a Angle) MilliRadian() Angle {
	aa := a.Radian()
	aa.Unit = MilliRadian
	aa.Value *= MilliRadianPerRadian
	return aa
}

// Degree returns an Angle in Degree units.
// No checks for NaN or +/-Inf
func (a Angle) Degree() Angle {
	ar := a.Radian()
	var aa Angle
	aa.Unit = Degree
	aa.Value = ar.Value / RadianPerDegree
	return aa
}

// Hour returns an Angle in Hour units.
// No checks for NaN or +/-Inf
func (a Angle) Hour() Angle {
	ar := a.Degree()
	var aa Angle
	aa.Unit = Hour
	aa.Value = ar.Value / DegreePerHour
	return aa
}

// ArcMinute returns an Angle in ArcMinutes units.
// No checks for NaN or +/-Inf
func (a Angle) ArcMinute() Angle {
	ad := a.Degree()
	ad.Unit = ArcMinute
	ad.Value *= ArcMinutePerDegree
	return ad
}

// ArcSecond returns an Angle in ArcSeconds units.
// No checks for NaN or +/-Inf
func (a Angle) ArcSecond() Angle {
	ad := a.Degree()
	ad.Unit = ArcSecond
	ad.Value *= ArcSecondPerDegree
	return ad
}

// MilliArcSecond returs an Angle in MilliArcSeconds units.
// No checks for NaN or +/-Inf
func (a Angle) MilliArcSecond() Angle {
	ad := a.Degree()
	ad.Unit = MilliArcSecond
	ad.Value *= MilliArcSecondPerDegree
	return ad
}

// Add sums the angles.
func (a Angle) Add(b Angle) Angle {
	v := a.Radian().Value + b.Radian().Value
	return NewAngle(Radian, v)
}

// Sub returns a - b.
func (a Angle) Sub(b Angle) Angle {
	v := a.Radian().Value - b.Radian().Value
	return NewAngle(Radian, v)
}

// Scale returns a scaled by b.
func (a Angle) Scale(b float64) Angle {
	v := a.Radian().Value * b
	return NewAngle(Radian, v)
}

// Normalize divides a by b.
func (a Angle) Normalize(b Angle) float64 {
	return a.Radian().Value / b.Radian().Value
}

// Div divides a by b.
func (a Angle) Div(b float64) (Angle, error) {
	var v Angle
	if b == 0.0 {
		return v, errors.New("Cannot divide by zero")
	}
	v = NewAngle(Radian, a.Radian().Value/b)
	return v, nil
}

// Abs returns the absolute value of 'a'
func (a Angle) Abs() Angle {
	return NewAngle(Radian, math.Abs(a.Radian().Value))
}

// Sign multiplies 'a' by the sign of b.
func (a Angle) Sign(b float64) Angle {
	absv := math.Abs(a.Radian().Value)
	sign := float64(1.0)
	if b < 0.0 {
		sign = -1.0
	}
	return NewAngle(Radian, absv*sign)
}

// Add sums the angles.  m is meant as a modulo function but
// can be any function with the same interface. Can be nil.
func (a Angle) AddModulo(b Angle, m func(float64) float64) Angle {
	//v := a.Degree().Value + b.Degree().Value
	v := a.Add(b).Radian().Value
	var vv float64
	if m != nil {
		vv = m(v)
	}
	return NewAngle(Radian, vv)
}

// Sub return the difference of the angles a - b.
// m is meant as a modulo function but
// can be any function with the same interface. Can be nil.
func (a Angle) SubModulo(b Angle, m func(float64) float64) Angle {
	//v := a.Degree().Value - b.Degree().Value
	v := a.Sub(b).Radian().Value
	var vv float64
	if m != nil {
		vv = m(v)
	}
	return NewAngle(Radian, vv)
}

// Constrain modifies reciever to be within the specified limits.
func (a *Angle) Constrain(mina, maxa Angle) {
	if a.Degree().Value > maxa.Degree().Value {
		//a.Value = maxa.Degree().Value
		*a = maxa
		return
	}
	if a.Degree().Value < mina.Degree().Value {
		//a.Value = mina.Degree().Value
		*a = mina
	}
}

// Map takes an angle and wrap count and  maps to a continous
// anglular range w/out discontinuities. For example -180-180 with a
// wrap count of +/-1 allows a mapping to -540-540 degrees.
// The 0 wrap starts a 0deg through 360(open).
func Map(a Angle, wrapCount int32) Angle {
	fac := float64(wrapCount) * DegreePerRevolution
	v := a.Degree().Value // [0-360)
	if wrapCount != 0 {
		v += fac
	}
	return NewAngle(Degree, v)
}

// Closest determines which wrap plane the desired angle should be on
// wc is the current wrap count and maxwc is the maximum allowed wrap count
func Closest(currentA, desAngle Angle, wc, maxwc int32) Angle {
	// don't allow angles outside maximum wrapcount
	wcp1 := wc + 1
	if wcp1 >= maxwc {
		wcp1 = wc
	}
	wcm1 := wc - 1
	if wcm1 <= -maxwc {
		wcm1 = wc
	}
	// map desired angle to current, plus 1 and minu 1 planes
	desAngleMap := Map(desAngle, wc)
	desAnglePlusOne := Map(desAngle, wcp1)
	desAngleMinusOne := Map(desAngle, wcm1)

	// subtract above angles from current angle extract the absolute value
	diff0 := currentA.Sub(desAngleMap)
	diffP1 := currentA.Sub(desAnglePlusOne)
	diffM1 := currentA.Sub(desAngleMinusOne)
	minDiffv := math.Abs(diff0.Degree().Value)
	diffPv := math.Abs(diffP1.Degree().Value)
	diffMv := math.Abs(diffM1.Degree().Value)

	// find the angle closet to the current angle and return it
	if diffMv < minDiffv {
		minDiffv = diffMv
		desAngleMap = desAngleMinusOne
	}
	if diffPv < minDiffv {
		desAngleMap = desAnglePlusOne
	}
	return desAngleMap
}

// Sin returns the sine of the angle
func (a Angle) Sin() float64 {
	return math.Sin(a.Radian().Value)
}

// Cos returns the cosine of the angle
func (a Angle) Cos() float64 {
	return math.Cos(a.Radian().Value)
}

// Tan returns the tangent of the angle
func (a Angle) Tan() float64 {
	return math.Tan(a.Radian().Value)
}

func (a Angle) GreaterThan(b Angle) bool {
	return a.Degree().Value > b.Degree().Value
}

func (a Angle) GreaterThanEqual(b Angle) bool {
	return a.Degree().Value >= b.Degree().Value
}

func (a Angle) LessThan(b Angle) bool {
	return a.Degree().Value < b.Degree().Value
}

func (a Angle) LessThanEqual(b Angle) bool {
	return a.Degree().Value <= b.Degree().Value
}

func (a Angle) Equal(b Angle) bool {
	return a.Degree().Value == b.Degree().Value
}

func (a Angle) Pow(f float64) Angle {
	return NewAngle(Radian, math.Pow(a.Radian().Value, f))
}
