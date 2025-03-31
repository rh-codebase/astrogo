package astrounit

// Length object

type LengthUnit int

const (
	// Enums to identify Length type
	none LengthUnit = iota
	Meter
	Centimeter
	Millimeter
	Micrometer
	Nanometer
	Femtometer
	Decimeter
	Kilometer

	// Length unit strings
	MeterStr      = "m"
	CentimeterStr = "cm"
	MillimeterStr = "mm"
	MicrometerStr = "um"
	NanometerStr  = "nm"
	FemtometerStr = "fm"
	DecimeterStr  = "dm"
	KilometerStr  = "km"
)

// Length supports a value and associated unit
type Length struct {
	Unit  LengthUnit `yaml:"unit"`
	Value float64    `yaml:"value"`
}

// NewLength returns a Length in the specified units and value
func NewLength(lu LengthUnit, v float64) Length {
	var l Length
	l.Unit = lu
	l.Value = v
	return l
}

// UnitString returns the units as a string for the given Length
func (l Length) UnitString() string {
	var ls string
	switch l.Unit {
	case Meter:
		ls = MeterStr
	case Centimeter:
		ls = CentimeterStr
	case Millimeter:
		ls = MillimeterStr
	case Nanometer:
		ls = NanometerStr
	case Femtometer:
		ls = FemtometerStr
	case Kilometer:
		ls = KilometerStr
	}
	return ls
}

func (l Length) Meter() Length {
	var ll Length
	ll.Unit = Meter
	switch l.Unit {
	case Meter:
		ll.Value = l.Value
	case Centimeter:
		ll.Value = l.Value * 1e-2
	case Millimeter:
		ll.Value = l.Value * 1e-3
	case Micrometer:
		ll.Value = l.Value * 1e-6
	case Nanometer:
		ll.Value = l.Value * 1e-9
	case Decimeter:
		ll.Value = l.Value * 1e2
	case Kilometer:
		ll.Value = l.Value * 1e3
	}
	return ll
}

func (l Length) Centimeter() Length {
	var ll Length
	ll.Unit = Centimeter
	ll.Value = l.Meter().Value * HectoF
	return ll
}

func (l Length) Millimeter() Length {
	var ll Length
	ll.Unit = Millimeter
	ll.Value = l.Meter().Value * KiloF
	return ll
}

func (l Length) Micrometer() Length {
	var ll Length
	ll.Unit = Micrometer
	ll.Value = l.Meter().Value * MegaF
	return ll
}

func (l Length) Nanometer() Length {
	var ll Length
	ll.Unit = Nanometer
	ll.Value = l.Meter().Value * GigaF
	return ll
}

func (l Length) Femtometer() Length {
	var ll Length
	ll.Unit = Femtometer
	ll.Value = l.Meter().Value * TeraF
	return ll
}

func (l Length) Decimeter() Length {
	var ll Length
	ll.Unit = Decimeter
	ll.Value = l.Meter().Value * DekaF
	return ll
}

func (l Length) Kilometer() Length {
	var ll Length
	ll.Unit = Kilometer
	ll.Value = l.Meter().Value * MilliF
	return ll
}
