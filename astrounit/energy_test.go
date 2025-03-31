package astrounit

import (
	"fmt"
	"testing"

	"github.com/rh-codebase/astrogo/testhelpers"
)

func TestNewEnergy(t *testing.T) {
	e, err := NewEnergy(Joule, Kilo, 12.23)
	if err != nil {
		fmt.Println("Error NewEnergy. Expected nil, Got ", err)
		t.Fail()
	}
	testhelpers.CheckF(t, e.Value, 12.23, "Value Error")
	testhelpers.CheckS(t, e.UnitString(), "kJ", "Unit Error")
	//fmt.Println(e.Value, " ", e.UnitString())
}

func TestConvert(t *testing.T) {
	e, _ := NewEnergy(Joule, Kilo, 1.0)
	fmt.Println(e.Value, "", e.UnitString())
	ee := e.Convert(Milli)
	fmt.Println(ee.Value, "", ee.UnitString())

	e, _ = NewEnergy(Joule, Milli, 1.0)
	fmt.Println(e.Value, "", e.UnitString())
	fmt.Println(e.Convert(Kilo).Value, "", e.Convert(Kilo).UnitString())
}

func TestSwitchUnit(t *testing.T) {
	e, _ := NewEnergy(Joule, Kilo, 1.e-19)
	ee := e.Electronvolt()
	fmt.Println(ee.Value, "", ee.UnitString())

	e, _ = NewEnergy(Electronvolt, Yotta, 1.)
	fmt.Println(e.Joule().Value, "", e.Joule().UnitString())
	fmt.Println(e.Joule().Convert(Kilo).Value, "", e.Joule().Convert(Kilo).UnitString())
}

func TestConvertTo(t *testing.T) {
	e, _ := NewEnergy(Joule, Kilo, 1.e-19)
	ee := e.ConvertTo(Electronvolt)
	fmt.Println(ee.Value, "", ee.UnitString())

	e, _ = NewEnergy(Electronvolt, Yotta, 1.)
	fmt.Println(e.ConvertTo(Joule).Value, "", e.ConvertTo(Joule).UnitString())
	fmt.Println(e.ConvertTo(Joule).Convert(Kilo).Value, "", e.ConvertTo(Joule).Convert(Kilo).UnitString())
}
