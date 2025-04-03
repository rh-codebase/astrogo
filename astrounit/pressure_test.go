package astrounit

import (
	"testing"

	th "github.com/rh-codebase/genutilsgo"
)

func TestToPascal(t *testing.T) {
	// map[Input] Expected
	cases := make(map[Pressure]Pressure)
	cases[NewPressure(Pascal, 10.1)] = NewPressure(Pascal, 10.1)
	cases[NewPressure(MilliPascal, 10.2)] = NewPressure(Pascal, 0.010199999999999999)
	cases[NewPressure(Millibar, 10.3)] = NewPressure(Pascal, .10300000000000001)
	cases[NewPressure(MillimeterHg, 10.4)] = NewPressure(Pascal, .07800663056359791)

	for ip, ep := range cases {
		gp := ip.ToPascal()
		th.CheckS(t, gp.UnitString(), ep.UnitString(), "Unit Error:")
		th.CheckF(t, gp.Value, ep.Value, "Value Error:")
	}

}

func TestToMilliPascal(t *testing.T) {
	// map[Input] Expected
	cases := make(map[Pressure]Pressure)
	cases[NewPressure(Pascal, 10.1)] = NewPressure(MilliPascal, 10100.0)
	cases[NewPressure(MilliPascal, 10.2)] = NewPressure(MilliPascal, 10.2)
	cases[NewPressure(Millibar, 10.3)] = NewPressure(MilliPascal, 103.00000000000001)
	cases[NewPressure(MillimeterHg, 10.4)] = NewPressure(MilliPascal, 78.00663056359791)

	for ip, ep := range cases {
		gp := ip.ToMilliPascal()
		th.CheckS(t, gp.UnitString(), ep.UnitString(), "Unit Error:")
		th.CheckF(t, gp.Value, ep.Value, "Value Error:")
	}
}

func TestToMillibar(t *testing.T) {
	// map[Input] Expected
	cases := make(map[Pressure]Pressure)
	cases[NewPressure(Pascal, 10.1)] = NewPressure(Millibar, 1010.0)
	cases[NewPressure(MilliPascal, 10.2)] = NewPressure(Millibar, 1.0199999999999998)
	cases[NewPressure(Millibar, 10.3)] = NewPressure(Millibar, 10.3)
	cases[NewPressure(MillimeterHg, 10.4)] = NewPressure(Millibar, 7.800663056359792)

	for ip, ep := range cases {
		gp := ip.ToMillibar()
		th.CheckS(t, gp.UnitString(), ep.UnitString(), "Unit Error:")
		th.CheckF(t, gp.Value, ep.Value, "Value Error:")
	}
}

func TestToMillimeterHg(t *testing.T) {
	// map[Input] Expected
	cases := make(map[Pressure]Pressure)
	cases[NewPressure(Pascal, 10.1)] = NewPressure(MillimeterHg, 1346.5522)
	cases[NewPressure(MilliPascal, 10.2)] = NewPressure(MillimeterHg, 1.3598843999999999)
	cases[NewPressure(Millibar, 10.3)] = NewPressure(MillimeterHg, 13.732166000000001)
	cases[NewPressure(MillimeterHg, 10.4)] = NewPressure(MillimeterHg, 10.4)

	for ip, ep := range cases {
		gp := ip.ToMillimeterHg()
		th.CheckS(t, gp.UnitString(), ep.UnitString(), "Unit Error:")
		th.CheckF(t, gp.Value, ep.Value, "Value Error:")
	}
}
