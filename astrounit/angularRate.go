// AngularRate type
package astrounit

import (
	"math"
	"time"
)

type AngularRateUnit int

const (
	// Enums to identify AngularRate type
	RadianPerSec AngularRateUnit = iota
	DegreePerSec
	MilliRadianPerSec
	ArcMinutePerSec
	ArcSecondPerSec
	MilliArcSecondPerSec
	HourPerSec

	// Angle unit strings
	RadianRateStr         = "rad/sec"
	DegreeRateStr         = "deg/sec"
	MilliRadianRateStr    = "mrad/sec"
	ArcMinuteRateStr      = "arcmin/sec"
	ArcSecondRateStr      = "arcsec/sec"
	MilliArcSecondRateStr = "mas/sec"
	HourRateStr           = "hr/sec"
)

// AngularRate supports a value and associated unit.
// always in per sec
type AngularRate struct {
	Unit     AngularRateUnit `yaml:"unit"`
	Angle    Angle           `yaml:"angle"`
	Timebase time.Duration   `yaml:"timebase"`
	Value    float64         `yaml:"value"`
}

// NewAngularRate returns an AngularRate from the
// specified units and value normalized to unit seconds.
func NewAngularRate(a Angle, t time.Duration) AngularRate {
	var ar AngularRate
	factor := float64(t) / float64(time.Second)
	ar.Unit = RadianPerSec
	ar.Value = a.Radian().Value / factor
	ar.Timebase = time.Second
	return ar
}

// UnitString returns the units as a string for the given AngularRate
func (a AngularRate) UnitString() string {
	var as string
	switch a.Unit {
	case RadianPerSec:
		as = RadianRateStr
	case DegreePerSec:
		as = DegreeRateStr
	case MilliRadianPerSec:
		as = MilliRadianRateStr
	case ArcMinutePerSec:
		as = ArcMinuteRateStr
	case ArcSecondPerSec:
		as = ArcSecondRateStr
	case MilliArcSecondPerSec:
		as = MilliArcSecondRateStr
	case HourPerSec:
		as = HourRateStr
	}
	return as
}

// Radian return an AngularRate in Radian/sec units.
// No checks for NaN or +/-Inf
func (a *AngularRate) Radian() AngularRate {
	var aa AngularRate
	aa.Unit = RadianPerSec
	switch a.Unit {
	case RadianPerSec:
		aa.Value = a.Value
	case MilliRadianPerSec:
		aa.Value = a.Value / MilliRadianPerRadian
	case DegreePerSec:
		aa.Value = a.Value * RadianPerDegree
	case ArcMinutePerSec:
		aa.Value = a.Value / ArcMinutePerDegree * RadianPerDegree
	case ArcSecondPerSec:
		aa.Value = a.Value / ArcSecondPerDegree * RadianPerDegree
	case MilliArcSecondPerSec:
		aa.Value = a.Value / ArcSecondPerDegree * RadianPerDegree / 1000.
	case HourPerSec:
		aa.Value = a.Value * DegreePerHour * RadianPerDegree
	}
	return aa
}

// MilliRadian returns an AngularRate in MilliRadian/sec units.
// No checks for NaN or +/-Inf
func (a *AngularRate) MilliRadian() AngularRate {
	aa := a.Radian()
	aa.Unit = MilliRadianPerSec
	aa.Value *= MilliRadianPerRadian
	return aa
}

// Degree returns an AngularRate in Degree/sec units.
// No checks for NaN or +/-Inf
func (a *AngularRate) Degree() AngularRate {
	ar := a.Radian()
	var aa AngularRate
	aa.Unit = DegreePerSec
	aa.Value = ar.Value / RadianPerDegree
	return aa
}

// Hour returns an AngularRate in Hour/sec units.
// No checks for NaN or +/-Inf
func (a *AngularRate) Hour() AngularRate {
	ar := a.Degree()
	var aa AngularRate
	aa.Unit = HourPerSec
	aa.Value = ar.Value / DegreePerHour
	return aa
}

// ArcMinute returns an AngularRate in ArcMinutes/sec units.
// No checks for NaN or +/-Inf
func (a *AngularRate) ArcMinute() AngularRate {
	ad := a.Degree()
	ad.Unit = ArcMinutePerSec
	ad.Value *= ArcMinutePerDegree
	return ad
}

// ArcSecond returns an AngularRate in ArcSeconds units.
// No checks for NaN or +/-Inf
func (a *AngularRate) ArcSecond() AngularRate {
	ad := a.Degree()
	ad.Unit = ArcSecondPerSec
	ad.Value *= ArcSecondPerDegree
	return ad
}

// MilliArcSecond returs an AngularRate in MilliArcSeconds units.
// No checks for NaN or +/-Inf
func (a *AngularRate) MilliArcSecond() AngularRate {
	ad := a.Degree()
	ad.Unit = MilliArcSecondPerSec
	ad.Value *= MilliArcSecondPerDegree
	return ad
}

// MultTime converts an Angular rate to an angle through multiplication by t[sec].
func (a *AngularRate) MultTime(t float64) Angle {
	val := a.Radian().Value * t // rad/sec * t[sec]-> rad
	return NewAngle(Radian, val)
}

// Sub returns the difference: 'a' - 'b'
func (a AngularRate) Sub(b AngularRate) AngularRate {
	v := NewAngle(Radian, a.Radian().Value-b.Radian().Value)
	return NewAngularRate(v, time.Second)
}

// Scale multiplies 'a' by a dimensionless factor
func (a AngularRate) Scale(f float64) AngularRate {
	nv := NewAngle(Radian, a.Radian().Value*f)
	return NewAngularRate(nv, time.Second)
}

// Normalize divides a by b and returns a  dimensionless number.
func (a AngularRate) Normalize(b AngularRate) float64 {
	return a.Radian().Value / b.Radian().Value
}

// Sign converts the sign of 'a' to match the sign of b.
func (a AngularRate) Sign(b float64) AngularRate {
	// Note: b does not have to be unity. Only the sign matters
	absv := math.Abs(a.Radian().Value)
	sign := float64(1.0)
	if b < 0.0 {
		sign = -1.0
	}
	nv := NewAngle(Radian, absv*sign)
	return NewAngularRate(nv, time.Second)
}

// Abs returns the absolute value of 'a'
func (a AngularRate) Abs() AngularRate {
	nv := NewAngle(Radian, math.Abs(a.Radian().Value))
	return NewAngularRate(nv, time.Second)
}

// GreaterThan returns true if AngularRate 'a' is > 'b'.
func (a AngularRate) GreaterThan(b AngularRate) bool {
	return a.Angle.GreaterThan(b.Angle)
}

// GreaterThanEqual returns true if AngularRate 'a' is >= 'b'.
func (a AngularRate) GreaterThanEqual(b AngularRate) bool {
	return a.Angle.GreaterThanEqual(b.Angle)
}

// LessThan returns true if AngularRate 'a' is < 'b'.
func (a AngularRate) LessThan(b AngularRate) bool {
	return a.Angle.LessThan(b.Angle)
}

// LessThanEqual returns true if AngularRate 'a' is <= to 'b'.
func (a AngularRate) LessThanEqual(b AngularRate) bool {
	return a.Angle.LessThanEqual(b.Angle)
}

// Equal returns true if AngularRate 'a' is equal to 'b'.
func (a AngularRate) Equal(b AngularRate) bool {
	return a.Angle.Equal(b.Angle)
}
