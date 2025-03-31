package astrounit

import (
	"math"
	"testing"
	"time"

	th "github.com/rh-codebase/astrogo/testhelpers"
)

const (
	artol = 1e-9
)

func TestNewAngularRate(t *testing.T) {
	a := NewAngle(Radian, .5)
	exp := 0.05

	tt := time.Duration(10.) * time.Second
	ar := NewAngularRate(a, tt)
	th.CheckFT(t, ar.Value, exp, artol, "Value Error")
	th.CheckS(t, ar.UnitString(), RadianRateStr, "Unit Error")
	//fmt.Println(ar.Value, ar.UnitString())
	//fmt.Println(ar.Angle.UnitString(), ar.Timebase.String())
	a = NewAngle(Degree, 90.)
	exp = math.Pi / 20.0
	tt = time.Duration(10.) * time.Second
	ar = NewAngularRate(a, tt)
	th.CheckFT(t, ar.Value, exp, artol, "Value Error")
	th.CheckS(t, ar.UnitString(), RadianRateStr, "Unit Error")
}

func TestARDegree(t *testing.T) {
	a := NewAngle(Radian, math.Pi/2.0)
	exp := 90.0
	tt := time.Second
	ar := NewAngularRate(a, tt)
	ard := ar.Degree()
	th.CheckFT(t, ard.Value, exp, artol, "Value Error")
	th.CheckS(t, ard.UnitString(), DegreeRateStr, "Unit Error")

	a = NewAngle(Radian, math.Pi/2.0)
	exp = 9.0
	tt = 10 * time.Second
	ar = NewAngularRate(a, tt)
	ard = ar.Degree()
	th.CheckFT(t, ard.Value, exp, artol, "Value Error")
	th.CheckS(t, ard.UnitString(), DegreeRateStr, "Unit Error")

	a = NewAngle(Radian, math.Pi/2.0)
	exp = 9000.0
	tt = 10 * time.Millisecond
	ar = NewAngularRate(a, tt)
	ard = ar.Degree()
	th.CheckFT(t, ard.Value, exp, artol, "Value Error")
	th.CheckS(t, ard.UnitString(), DegreeRateStr, "Unit Error")

}

func TestMultTime(t *testing.T) {
	a := NewAngle(Radian, 0.5)
	ar := NewAngularRate(a, time.Second)
	d := 0.100
	nr := ar.MultTime(d)
	th.CheckFT(t, 0.05, nr.Radian().Value, artol, "Error Value:")
}
