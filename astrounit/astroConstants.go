/*
* package astro contains constants and functions related to astronomy
 */
package astrounit

import (
	"errors"
	"math"
)

const (
	// SI units
	QuettaF = float64(1.e30) // Q
	RonnaF  = float64(1.e27) // R
	YottaF  = float64(1.e24) // Y
	ZettaF  = float64(1.e21) // Z
	ExaF    = float64(1.e18) // E
	PetaF   = float64(1.e15) // P
	TeraF   = float64(1.e12) // T
	GigaF   = float64(1.e9)  // G
	MegaF   = float64(1.e6)  // M
	KiloF   = float64(1.e3)  // k
	HectoF  = float64(1.e2)  // h
	DekaF   = float64(1.e1)  // da
	OneF    = 1.0
	DeciF   = float64(1.e-1)  // d
	CentiF  = float64(1.e-2)  // c
	MilliF  = float64(1.e-3)  // m
	MicroF  = float64(1.e-6)  // u
	NanoF   = float64(1.e-9)  // n
	PicoF   = float64(1.e-12) // p
	FemtoF  = float64(1.e-15) // f
	AttoF   = float64(1.e-18) // a
	ZeptoF  = float64(1.e-21) // z
	YoctoF  = float64(1.e-24) // y
	RontoF  = float64(1.e-27) // r
	QuectoF = float64(1.e-30) // q

	// SI Symbols
	// TODO: store in slice using SIPrefix as index
	QuettaS = "Q"
	RonnaS  = "R"
	YottaS  = "Y"
	ZettaS  = "Z"
	ExaS    = "E"
	PetaS   = "P"
	TeraS   = "T"
	GigaS   = "G"
	MegaS   = "M"
	KiloS   = "k"
	HectoS  = "h"
	DekaS   = "da"
	OneS    = ""
	DeciS   = "d"
	CentiS  = "c"
	MilliS  = "m"
	MicroS  = "u"
	NanoS   = "n"
	PicoS   = "p"
	FemtoS  = "f"
	AttoS   = "a"
	ZeptoS  = "z"
	YoctoS  = "y"
	RontoS  = "r"
	QuectoS = "q"

	// SI Names
	QuettaN = "Q"
	RonnaN  = "R"
	YottaN  = "Y"
	ZettaN  = "Z"
	ExaN    = "E"
	PetaN   = "P"
	TeraN   = "T"
	GigaN   = "G"
	MegaN   = "million"
	KiloN   = "thousand"
	HectoN  = "h"
	DekaN   = "da"
	OneN    = ""
	DeciN   = "d"
	CentiN  = "c"
	MilliN  = "thousandths"
	MicroN  = "u"
	NanoN   = "n"
	PicoN   = "p"
	FemtoN  = "f"
	AttoN   = "a"
	ZeptoN  = "z"
	YoctoN  = "y"
	RontoN  = "r"
	QuectoN = "q"

	// zenith
	ZenithDeg = float64(90.0)

	// conversion factor from Astronomical Almanac
	SolarDay    = float64(0.99726956633)
	SiderealDay = 1.0 / SolarDay
	// rates
	MinutePerDegree = float64(60.0)
	SecondPerDegree = float64(3600.0)
	HourPerRadian   = float64(12.0 / math.Pi)
	//	HourPerSiderealDay         = HourPerDay *
	RadianPerDegree            = float64(math.Pi / 180.)
	ArcMinutePerDegree         = float64(60.0)
	ArcSecondPerDegree         = float64(3600.0)
	MilliArcSecondPerDegree    = float64(3600000.0)
	RadianPerMinute            = float64(math.Pi / (180. * 60.))
	RadianPerSecond            = float64(math.Pi / (180. * 3600.))
	MilliRadianPerRadian       = KiloF
	MilliArcSecondPerArcSecond = KiloF
	DegreePerHour              = float64(15.0)
	DegreePerRevolution        = float64(360.0)
	// 3 Coefficients of Clausius-Clapeyron equation approximation
	// as derived by Crane (1976). See, e.g. TMS equation 13.15
	CC_A                  = float64(6.105)   // leading coeff.
	CC_B                  = float64(25.22)   // T coeff.
	CC_C                  = float64(5.31)    // ln(T) coeff.
	AbsoluteZeroCelsius   = float64(-273.15) // Celsius
	MillimeterHgPerPascal = float64(133.322)
	MillibarPerPascal     = HectoF
	MilliPascalPerPascal  = KiloF
	FahrenheitPerCelsius  = float64(1.8)
	WaterFreezeFahrenheit = float64(32.0)
	MilliKelvinPerKelvin  = KiloF
	JoulePerElectronvolt  = float64(1.602176634e-19)
	JoulePerErg           = float64(1.0e-7)
	R_Water               = 461.5 //Gas Const for Water. Units: Joule/(kg*K)
	Rho_Water             = 1.0   // water density. Units: g/cm^3

	/**
	 * Coefficients for modified Clausius-Clapeyron equations
	 * used in NOAA's Advanced Weather Interactive Processing System
	 * Modified Clausius-Clapeyron is shown here (but not derived!):
	 * http://meted.ucar.edu/awips/validate/dewpnt.htm
	 */
	AWIPS_C3  = 223.1986
	AWIPS_C4  = 0.0182758048
	AWIPS_C15 = 26.66082

	/**
	 * Coefficients of Magnus-Teton formula of dewpoint.
	 * Lawrence, Mark G. (1 February 2005). "The Relationship between Relative Humidity and the Dewpoint Temperature in Moist Air: A Simple Conversion and Applications". Bulletin of the American Meteorological Society. 86 (2): 225–234. Bibcode:2005BAMS...86..225L. doi:10.1175/BAMS-86-2-225. Retrieved 15 March 2024.
	 */
	MT_B = 17.625 // unit: Celsius
	MT_C = 243.04 // unit: Celsius

	/**
	 * Coefficients for Hoffman-Welch (BIMA) formulation
	 * of dewpoint. From HatCreek's weatherman1.c code.
	 */
	HW_A = 1.598e9
	HW_B = 5370

	/**
	 *  Earth radius
	 *  Units: m
	 */
	EARTH_RADIUS = 6.3781366e6 // from NOVAS
)

type SIPrefix int

const (
	_ SIPrefix = iota
	Quetta
	Ronna
	Yotta
	Zetta
	Exa
	Peta
	Tera
	Giga
	Mega
	Kilo
	Hecto
	Deka
	One
	Deci
	Centi
	Milli
	Micro
	Nano
	Pico
	Femto
	Atto
	Zepto
	Yocto
	Ronto
	Quecto
)

type SI struct {
	Enum   SIPrefix
	Symbol string
	Factor float64
	Name   string
}

func NewSI(sip SIPrefix) (SI, error) {
	var si SI
	switch sip {
	case Quetta:
		si.Enum = sip
		si.Symbol = QuettaS
		si.Factor = QuettaF
		si.Name = QuettaN
	case Ronna:
		si.Enum = sip
		si.Symbol = RonnaS
		si.Factor = RonnaF
		si.Name = RonnaN
	case Yotta:
		si.Enum = sip
		si.Symbol = YottaS
		si.Factor = YottaF
		si.Name = YottaN
	case Zetta:
		si.Enum = sip
		si.Symbol = ZettaS
		si.Factor = ZettaF
		si.Name = ZettaN
	case Exa:
		si.Enum = sip
		si.Symbol = ExaS
		si.Factor = ExaF
		si.Name = ExaN
	case Peta:
		si.Enum = sip
		si.Symbol = PetaS
		si.Factor = PetaF
		si.Name = PetaN
	case Tera:
		si.Enum = sip
		si.Symbol = TeraS
		si.Factor = TeraF
		si.Name = TeraN
	case Giga:
		si.Enum = sip
		si.Symbol = GigaS
		si.Factor = GigaF
		si.Name = GigaN
	case Mega:
		si.Enum = sip
		si.Symbol = MegaS
		si.Factor = MegaF
		si.Name = MegaN
	case Kilo:
		si.Enum = sip
		si.Symbol = KiloS
		si.Factor = KiloF
		si.Name = KiloN
	case Hecto:
		si.Enum = sip
		si.Symbol = HectoS
		si.Factor = HectoF
		si.Name = HectoN
	case Deka:
		si.Enum = sip
		si.Symbol = DekaS
		si.Factor = DekaF
		si.Name = DekaN
	case One:
		si.Enum = sip
		si.Symbol = OneS
		si.Factor = OneF
		si.Name = OneN
	case Deci:
		si.Enum = sip
		si.Symbol = DeciS
		si.Factor = DeciF
		si.Name = DeciN
	case Centi:
		si.Enum = sip
		si.Symbol = CentiS
		si.Factor = CentiF
		si.Name = CentiN
	case Milli:
		si.Enum = sip
		si.Symbol = MilliS
		si.Factor = MilliF
		si.Name = MilliN
	case Micro:
		si.Enum = sip
		si.Symbol = MicroS
		si.Factor = MicroF
		si.Name = MicroN
	case Nano:
		si.Enum = sip
		si.Symbol = NanoS
		si.Factor = NanoF
		si.Name = NanoN
	case Pico:
		si.Enum = sip
		si.Symbol = PicoS
		si.Factor = PicoF
		si.Name = PicoN
	case Femto:
		si.Enum = sip
		si.Symbol = FemtoS
		si.Factor = FemtoF
		si.Name = FemtoN
	case Atto:
		si.Enum = sip
		si.Symbol = AttoS
		si.Factor = AttoF
		si.Name = AttoN
	case Zepto:
		si.Enum = sip
		si.Symbol = ZeptoS
		si.Factor = ZeptoF
		si.Name = ZeptoN
	case Yocto:
		si.Enum = sip
		si.Symbol = YoctoS
		si.Factor = YoctoF
		si.Name = YoctoN
	case Ronto:
		si.Enum = sip
		si.Symbol = RontoS
		si.Factor = RontoF
		si.Name = RontoN
	case Quecto:
		si.Enum = sip
		si.Symbol = QuectoS
		si.Factor = QuectoF
		si.Name = QuectoN
	default:
		return si, errors.New("Uknown Prefix")
	}
	return si, nil
}

// SIPrefixes returns a slice of SIPrefix. Primarily used for UI dropdown
// creation.
func SIPrefixes() []SIPrefix {
	sip := make([]SIPrefix, 25)
	sip[0] = Quetta
	sip[1] = Ronna
	sip[2] = Yotta
	sip[3] = Zetta
	sip[4] = Exa
	sip[5] = Peta
	sip[6] = Tera
	sip[7] = Giga
	sip[8] = Mega
	sip[9] = Kilo
	sip[10] = Hecto
	sip[11] = Deka
	sip[12] = One
	sip[13] = Deci
	sip[14] = Centi
	sip[15] = Milli
	sip[16] = Micro
	sip[17] = Nano
	sip[18] = Pico
	sip[19] = Femto
	sip[20] = Atto
	sip[21] = Zepto
	sip[22] = Yocto
	sip[23] = Ronto
	sip[24] = Quecto
	return sip
}

type ConformalQuantity interface {
	UnitString() string
}

func UnitString(cq ConformalQuantity) string {
	return cq.UnitString()
}
