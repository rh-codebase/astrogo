package astrounit

import (
	"math"
	"testing"

	th "github.com/rh-codebase/astrogo/testhelpers"
)

func TestNewTemperature(t *testing.T) {
	_, err := NewTemperature(Kelvin, -1.0)
	th.CheckErrorNil(t, err, "Error:")

	_, err = NewTemperature(Kelvin, math.NaN())
	th.CheckErrorNil(t, err, "Error:")

	_, err = NewTemperature(Kelvin, math.Inf(1))
	th.CheckErrorNil(t, err, "Error:")

	_, err = NewTemperature(Kelvin, math.Inf(-1))
	th.CheckErrorNil(t, err, "Error:")

	it, err := NewTemperature(Kelvin, 10.1)
	th.CheckError(t, err, nil, "nil Error:")
	th.CheckS(t, it.UnitString(), KelvinStr, "Unit Error")
	th.CheckF(t, it.Value, 10.1, "Value Error")

	it, err = NewTemperature(MilliKelvin, 10.1)
	th.CheckError(t, err, nil, "nil Error:")
	th.CheckS(t, it.UnitString(), MilliKelvinStr, "Unit Error")
	th.CheckF(t, it.Value, 10.1, "Value Error")

	it, err = NewTemperature(Celsius, 10.1)
	th.CheckError(t, err, nil, "nil Error:")
	th.CheckS(t, it.UnitString(), CelsiusStr, "Unit Error")
	th.CheckF(t, it.Value, 10.1, "Value Error")

	it, _ = NewTemperature(Fahrenheit, 10.1)
	th.CheckError(t, err, nil, "nil Error:")
	th.CheckS(t, it.UnitString(), FahrenheitStr, "Unit Error")
	th.CheckF(t, it.Value, 10.1, "Value Error")

	it, _ = NewTemperature(Fahrenheit, -10.1)
	th.CheckError(t, err, nil, "nil Error:")
	th.CheckS(t, it.UnitString(), FahrenheitStr, "Unit Error")
	th.CheckF(t, it.Value, -10.1, "Value Error")
}

func TestToKelvin(t *testing.T) {
	cases := make(map[Temperature]Temperature)
	it, _ := NewTemperature(Kelvin, 10.1)
	et, _ := NewTemperature(Kelvin, 10.1)
	cases[it] = et
	it, _ = NewTemperature(MilliKelvin, 10.2)
	et, _ = NewTemperature(Kelvin, 0.010199999999999999)
	cases[it] = et
	it, _ = NewTemperature(Celsius, 10.3)
	et, _ = NewTemperature(Kelvin, 283.45)
	cases[it] = et
	it, _ = NewTemperature(Fahrenheit, 32.0)
	et, _ = NewTemperature(Kelvin, 273.15)
	cases[it] = et

	for it, et = range cases {
		gt := it.ToKelvin()
		th.CheckS(t, gt.UnitString(), et.UnitString(), "Unit Error")
		th.CheckFT(t, gt.Value, et.Value, 1e-10, "Value Error")
	}
}

func TestToMilliKelvin(t *testing.T) {
	cases := make(map[Temperature]Temperature)
	it, _ := NewTemperature(Kelvin, 10.1)
	et, _ := NewTemperature(MilliKelvin, 10100.0)
	cases[it] = et
	it, _ = NewTemperature(MilliKelvin, 10.2)
	et, _ = NewTemperature(MilliKelvin, 10.2)
	cases[it] = et
	it, _ = NewTemperature(Celsius, 10.3)
	et, _ = NewTemperature(MilliKelvin, 283450.0)
	cases[it] = et
	it, _ = NewTemperature(Fahrenheit, 32.0)
	et, _ = NewTemperature(MilliKelvin, 273150)
	cases[it] = et

	for it, et = range cases {
		gt := it.ToMilliKelvin()
		th.CheckS(t, gt.UnitString(), et.UnitString(), "Unit Error")
		th.CheckFT(t, gt.Value, et.Value, 1e-10, "Value Error")
	}
}

func TestToCelsius(t *testing.T) {
	cases := make(map[Temperature]Temperature)
	it, _ := NewTemperature(Kelvin, 10.1)
	et, _ := NewTemperature(Celsius, -263.04999999999995)
	cases[it] = et
	it, _ = NewTemperature(MilliKelvin, 1000.15)
	et, _ = NewTemperature(Celsius, -272.14984999999996)
	cases[it] = et
	it, _ = NewTemperature(Celsius, 10.3)
	et, _ = NewTemperature(Celsius, 10.3)
	cases[it] = et
	it, _ = NewTemperature(Fahrenheit, 32.0)
	et, _ = NewTemperature(Celsius, 0.0)
	cases[it] = et

	for it, et = range cases {
		gt := it.ToCelsius()
		th.CheckS(t, gt.UnitString(), et.UnitString(), "Unit Error")
		th.CheckFT(t, gt.Value, et.Value, 1e-10, "Value Error")
	}
}

func TestToFahrenheit(t *testing.T) {
	cases := make(map[Temperature]Temperature)
	it, _ := NewTemperature(Kelvin, 10.1)
	et, _ := NewTemperature(Fahrenheit, -441.48999999999995)
	cases[it] = et
	it, _ = NewTemperature(MilliKelvin, 1000.15)
	et, _ = NewTemperature(Fahrenheit, -457.86972999999995)
	cases[it] = et
	it, _ = NewTemperature(Celsius, 0.0)
	et, _ = NewTemperature(Fahrenheit, 32.0)
	cases[it] = et
	it, _ = NewTemperature(Fahrenheit, 32.0)
	et, _ = NewTemperature(Fahrenheit, 32.0)
	cases[it] = et

	for it, et = range cases {
		gt := it.ToFahrenheit()
		th.CheckS(t, gt.UnitString(), et.UnitString(), "Unit Error")
		th.CheckFT(t, gt.Value, et.Value, 1e-10, "Value Error")
	}
}
