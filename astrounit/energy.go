// Energy type
package astrounit

import (
	"errors"
	"fmt"
)

type EnergyUnit int

const (
	// Enums to identify Energy type
	_ EnergyUnit = iota
	Joule
	/*
		Decijoule
		Millijoule
		Microjoule
		Nanojoule
		Femptojoule
		Kilojoule
		Megajoule
		Gigajoule
	*/
	Erg
	Electronvolt
	/*
		MegaElectronjoule
		GigaElectronjoule
	*/

	// Energy unit strings
	JouleStr            = "J"
	DecijouleStr        = "dJ"
	CentijouleStr       = "cJ"
	MillijouleStr       = "mJ"
	MicrojouleStr       = "uJ"
	NanojouleStr        = "nJ"
	PicojouleStr        = "pJ"
	FemptojouleStr      = "fJ"
	DekajouleStr        = "daJ"
	HectojouleStr       = "hJ"
	KilojouleStr        = "kJ"
	MegajouleStr        = "MJ"
	GigajouleStr        = "GJ"
	ElectronvoltStr     = "eV"
	KiloElectronvoltStr = "KeV"
	MegaElectronvoltStr = "MeV"
	GigaElectronvoltStr = "GeV"
	ErgStr              = "erg"
)

// Energy supports a value and associated unit
type Energy struct {
	Unit  EnergyUnit `yaml:"unit"`
	Si    SI
	Value float64 `yaml:"value"`
}

func checkEnergyUnit(u EnergyUnit) error {
	switch u {
	case Joule, Electronvolt, Erg:
		return nil
	default:
		emsg := fmt.Sprintf("Unknown Energy unit: %d", u)
		return errors.New(emsg)
	}
}

// NewEnergy returns a Energy in the specified units and value. Return non nil
// error if EnergyUnit is unknown.
func NewEnergy(eu EnergyUnit, sip SIPrefix, v float64) (Energy, error) {
	var e Energy
	if err := checkEnergyUnit(eu); err != nil {
		return e, err
	}
	e.Unit = eu
	si, err := NewSI(sip)
	if err != nil {
		return e, err
	}
	e.Si = si
	e.Value = v
	return e, nil
}

// EnergyUnits returns a slice of EnergyUnit. Primarily used for
// UI dropdown creation.
func EnergyUnits() []EnergyUnit {
	eu := make([]EnergyUnit, 3)
	eu[0] = Joule
	eu[1] = Erg
	eu[2] = Electronvolt
	return eu
}

// UnitString returns the units as a string for the given Energy
func (e Energy) UnitString() string {
	var es string
	switch e.Unit {
	case Joule:
		es = JouleStr
	case Erg:
		es = ErgStr
	case Electronvolt:
		es = ElectronvoltStr
	}
	s := fmt.Sprintf("%s%s", e.Si.Symbol, es)

	return s
}

func (e Energy) Joule() Energy {
	var ee Energy
	ee.Unit = Joule
	ee.Si, _ = NewSI(One)
	switch e.Unit {
	case Joule:
		ee.Value = e.Value
	case Erg:
		ee.Value = e.Value * JoulePerErg
	case Electronvolt:
		ee.Value = e.Value * JoulePerElectronvolt
	}
	// Handle SI prefix
	ee.Value *= e.Si.Factor
	return ee
}

func (e Energy) Electronvolt() Energy {
	var ee Energy
	ee.Unit = Electronvolt
	ee.Si, _ = NewSI(One)
	ee.Value = e.Joule().Value / JoulePerElectronvolt
	return ee
}

// ConvertTo converts between the different Energy Units.
// This function works well with UI where a dropdown can specify
// EnergyUnit and then this function is called for display.
func (e Energy) ConvertTo(eu EnergyUnit) Energy {
	var ee Energy
	ee.Unit = eu
	ee.Si, _ = NewSI(One)
	switch eu {
	case Joule:
		ee.Value = e.Joule().Value
	case Erg:
		ee.Value = e.Joule().Value / JoulePerErg
	case Electronvolt:
		ee.Value = e.Joule().Value / JoulePerElectronvolt
	}
	return ee
}

func (e Energy) Convert(sip SIPrefix) Energy {
	var ee Energy
	ee.Unit = e.Unit
	ee.Si, _ = NewSI(sip)
	ee.Value = e.Value * e.Si.Factor / ee.Si.Factor
	return ee
}
