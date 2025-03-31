// Pressure type
package astrounit

type PressureUnit int

const (
	Pascal PressureUnit = iota
	Millibar
	MillimeterHg
	MilliPascal

	PascalStr       = "Pa"
	MillimeterHgStr = "mmHg"
	MillibarStr     = "mb"
	MilliPascalStr  = "mPa"
)

type Pressure struct {
	Unit  PressureUnit
	Value float64
}

func NewPressure(pu PressureUnit, v float64) Pressure {
	var p Pressure
	p.Unit = pu
	p.Value = v
	return p
}

func (p Pressure) UnitString() string {
	var us string
	switch p.Unit {
	case Pascal:
		us = PascalStr
	case MilliPascal:
		us = MilliPascalStr
	case Millibar:
		us = MillibarStr
	case MillimeterHg:
		us = MillimeterHgStr
	}
	return us
}

func (p Pressure) ToPascal() Pressure {
	var pp Pressure
	pp.Unit = Pascal
	switch p.Unit {
	case Pascal:
		pp.Value = p.Value
	case MilliPascal:
		pp.Value = p.Value / MilliPascalPerPascal
	case Millibar:
		pp.Value = p.Value / MillibarPerPascal
	case MillimeterHg:
		pp.Value = p.Value / MillimeterHgPerPascal
	}
	return pp
}

func (p Pressure) ToMilliPascal() Pressure {
	var pp Pressure
	pp.Unit = MilliPascal
	pp.Value = p.ToPascal().Value * 1000.0
	return pp
}

func (p Pressure) ToMillibar() Pressure {
	var pp Pressure
	pp.Unit = Millibar
	pp.Value = p.ToPascal().Value * MillibarPerPascal
	return pp
}

func (p Pressure) ToMillimeterHg() Pressure {
	var pp Pressure
	pp.Unit = MillimeterHg
	pp.Value = p.ToPascal().Value * MillimeterHgPerPascal
	return pp
}
