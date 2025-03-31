package astrounit

import "time"

type AngleCoord struct {
	A1 Angle
	A2 Angle
}

type AngleCoordEpoch struct {
	Epoch time.Time
	Ac    AngleCoord
}

func NewAzElCoord(au AngleUnit, az, el float64) AngleCoord {
	var a AngleCoord
	a.A1 = NewAngle(au, az)
	a.A2 = NewAngle(au, el)

	return a
}

func NewAzElCoordA(az, el Angle) AngleCoord {
	var a AngleCoord
	a.A1 = az
	a.A2 = el

	return a
}

func NewAzElCoordEpoch(au AngleUnit, az, el float64, epoch time.Time) AngleCoordEpoch {
	var ae AngleCoordEpoch
	ae.Epoch = epoch
	ae.Ac = NewAzElCoord(au, az, el)

	return ae
}

// NewRaDecCoord returns an AngleCoord representing an Ra and Dec. Since
// Ra is normally in units of hours and Dec in degrees, this function allows
// each unit to be specified.
func NewRaDecCoord(rau AngleUnit, ra float64,
	dau AngleUnit, dec float64) AngleCoord {
	var a AngleCoord
	a.A1 = NewAngle(rau, ra)
	a.A2 = NewAngle(dau, dec)

	return a
}

// NewRaDecCoord returns an AngleCoord representing an Ra and Dec. Since
// Ra is normally in units of hours and Dec in degrees, this function allows
// each unit to be specified.
func NewRaDecCoordA(ra, dec Angle) AngleCoord {
	var a AngleCoord
	a.A1 = ra
	a.A2 = dec

	return a
}

func NewRaDecCoordEpoch(rau AngleUnit, ra float64, dau AngleUnit, dec float64, epoch time.Time) AngleCoordEpoch {
	var ae AngleCoordEpoch
	ae.Epoch = epoch
	ae.Ac = NewRaDecCoord(rau, ra, dau, dec)

	return ae
}

// Receiver functions

func (ac AngleCoord) Az() Angle {
	return ac.A1
}

func (ac AngleCoord) El() Angle {
	return ac.A2
}

func (ac AngleCoord) Ra() Angle {
	return ac.A1
}

func (ac AngleCoord) Dec() Angle {
	return ac.A2
}

func (ae AngleCoordEpoch) Az() Angle {
	return ae.Ac.A1
}

func (ae AngleCoordEpoch) El() Angle {
	return ae.Ac.A2
}

func (ae AngleCoordEpoch) AzEl() AngleCoord {
	return ae.Ac
}

func (ae AngleCoordEpoch) Ra() Angle {
	return ae.Ac.A1
}

func (ae AngleCoordEpoch) Dec() Angle {
	return ae.Ac.A2
}

func (ae AngleCoordEpoch) RaDec() AngleCoord {
	return ae.Ac
}

// Lat, Long Coordinate
func NewLatLonCoord(au AngleUnit, lat, long float64) AngleCoord {
	var a AngleCoord
	a.A1 = NewAngle(au, lat)
	a.A2 = NewAngle(au, long)

	return a
}

func NewLatLonCoordA(lat, lon Angle) AngleCoord {
	var a AngleCoord
	a.A1 = lat
	a.A2 = lon

	return a
}

func (ac AngleCoord) Lat() Angle {
	return ac.A1
}

func (ac AngleCoord) Lon() Angle {
	return ac.A2
}
