package astrounit

import (
	"fmt"
	"math"
	"testing"

	th "github.com/rh-codebase/genutilsgo"
)

func TestWaterSaturatedPressure(t *testing.T) {
	cases := make(map[Temperature]Pressure)
	tt, _ := NewTemperature(Kelvin, 273.15)
	cases[tt] = NewPressure(Millibar, 6.105)
	tt, _ = NewTemperature(Celsius, 0.0)
	cases[tt] = NewPressure(Millibar, 6.105)
	tt, _ = NewTemperature(Celsius, 10.0)
	cases[tt] = NewPressure(Millibar, 12.291123450237006)
	tt, _ = NewTemperature(Celsius, -40.0)
	cases[tt] = NewPressure(Millibar, 0.1869452203246955)
	tt, _ = NewTemperature(Celsius, 40.0)
	cases[tt] = NewPressure(Millibar, 74.0615441352747)
	tol := 1e-15

	for in, ex := range cases {
		g, err := WaterSaturatedPressure(in)
		th.CheckError(t, err, nil, "Return Error")
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckFT(t, g.Value, ex.Value, tol, "Value Error")
	}

	// Feed in bad input and make sure and non nil error is returned.
	tt, _ = NewTemperature(Celsius, math.NaN())
	_ = NewPressure(Millibar, 74.0615441352747)
	_, err := WaterSaturatedPressure(tt)
	th.CheckErrorNil(t, err, "Return Error")

	tt, _ = NewTemperature(Celsius, math.Inf(1))
	_ = NewPressure(Millibar, 74.0615441352747)
	_, err = WaterSaturatedPressure(tt)
	th.CheckErrorNil(t, err, "Return Error")

	tt, _ = NewTemperature(Celsius, math.Inf(-1))
	_ = NewPressure(Millibar, 74.0615441352747)
	_, err = WaterSaturatedPressure(tt)
	th.CheckErrorNil(t, err, "Return Error")

}

func TestUlichMappingFunction(t *testing.T) {
	for idx := 0; idx < 90; idx++ {
		an := NewAngle(Degree, float64(idx))
		got := ulichMappingFunction(an)
		fmt.Println(got)
	}
}

func TestZenithRefractivity(t *testing.T) {
	airT, _ := NewTemperature(Celsius, 10.0)
	atmP := NewPressure(Millibar, 1013.25) // mb=hPa: 1013.25
	rh := 20.0
	freq := 1.e26

	zr, err := ZenithRefractivity(airT, atmP, rh, freq)
	if err != nil {
		fmt.Println("ZenithRefractivity retuned err: ", err)
		t.Fail()
	}
	fmt.Println(zr)

}

func TestComputeRefractionCorrection(t *testing.T) {
	airT, _ := NewTemperature(Celsius, 10.0)
	atmP := NewPressure(Millibar, 1013.25)
	rh := 50.
	freq := 1.e6
	alt := NewLength(Meter, 3.0)
	for idx := 0; idx <= 90; idx += 5 {
		elevation := NewAngle(Degree, float64(idx))
		ref, err := ComputeRefractionCorrection(airT, atmP, rh, elevation, freq, alt)
		if err != nil {
			fmt.Println("ComputeRefractionCorrection returned err: ", err)
			t.Fail()
		}
		fmt.Printf("Elevation[deg]: %d, Refraction Correction[arcmin]: %5.3f\n", idx, ref.ArcMinute().Value)
	}
}
