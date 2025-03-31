// Temperature type
package astrounit

import (
	"errors"
	"fmt"
	"math"
)

type TemperatureUnit int

const (
	Kelvin TemperatureUnit = iota
	MilliKelvin
	Celsius
	Fahrenheit

	KelvinStr      = "K"
	MilliKelvinStr = "mK"
	CelsiusStr     = "C"
	FahrenheitStr  = "F"
)

type Temperature struct {
	Unit  TemperatureUnit
	Value float64
}

// NewTemperature creates a Temperature type with given unit and value.
// Error on negative Kelvin and NaN, +/-Inf
func NewTemperature(tu TemperatureUnit, v float64) (Temperature, error) {
	var t Temperature
	if math.IsNaN(v) {
		emsg := fmt.Sprintf("Illegal temperature of NaN")
		return t, errors.New(emsg)
	}
	if math.IsInf(v, 1) {
		emsg := fmt.Sprintf("Illegal temperature of +Inf")
		return t, errors.New(emsg)
	}
	if math.IsInf(v, -1) {
		emsg := fmt.Sprintf("Illegal temperature of -Inf")
		return t, errors.New(emsg)
	}
	if tu == Kelvin && v < 0.0 {
		emsg := fmt.Sprintf("Illegal temperature: %f Kelvin", v)
		return t, errors.New(emsg)
	}
	t.Unit = tu
	t.Value = v
	return t, nil
}

func (t *Temperature) UnitString() string {
	var us string
	switch t.Unit {
	case Kelvin:
		us = KelvinStr
	case MilliKelvin:
		us = MilliKelvinStr
	case Celsius:
		us = CelsiusStr
	case Fahrenheit:
		us = FahrenheitStr
	}
	return us
}

// ftoc converters Fahrenheit to Celsius
func ftoc(f float64) float64 {
	return (f - WaterFreezeFahrenheit) / FahrenheitPerCelsius
}

// ctof converters Celsius to Fahrenheit
func ctof(c float64) float64 {
	return c*FahrenheitPerCelsius + WaterFreezeFahrenheit
}

// ToKelvin return the temperaure in Kelvin
func (t *Temperature) ToKelvin() Temperature {
	var tt Temperature
	tt.Unit = Kelvin
	switch t.Unit {
	case Kelvin:
		tt.Value = t.Value
	case MilliKelvin:
		tt.Value = t.Value / MilliKelvinPerKelvin
	case Celsius:
		tt.Value = t.Value - AbsoluteZeroCelsius
	case Fahrenheit:
		tt.Value = ftoc(t.Value) - AbsoluteZeroCelsius
	}
	return tt
}

// ToMilliKelvin return the temperaure in Kelvin
func (t *Temperature) ToMilliKelvin() Temperature {
	t0 := t.ToKelvin()
	var tt Temperature
	tt.Unit = MilliKelvin
	tt.Value = t0.Value * MilliKelvinPerKelvin
	return tt
}

// ToFahrenheit returns the temperature in Fahrenheit
func (t *Temperature) ToFahrenheit() Temperature {
	var tt Temperature
	tt.Unit = Fahrenheit
	switch t.Unit {
	case Fahrenheit:
		tt.Value = t.Value
	case Kelvin:
		c := t.Value + AbsoluteZeroCelsius
		tt.Value = ctof(c)
	case MilliKelvin:
		c := t.Value/MilliKelvinPerKelvin + AbsoluteZeroCelsius
		tt.Value = ctof(c)
	case Celsius:
		tt.Value = ctof(t.Value)
	}
	return tt
}

// ToCelsius returns the temperature in Celsius
func (t *Temperature) ToCelsius() Temperature {
	var tt Temperature
	tt.Unit = Celsius
	if t.Unit == Celsius {
		tt = *t
		return tt
	}
	t0 := t.ToKelvin()

	tt.Value = t0.Value + AbsoluteZeroCelsius
	return tt
}
