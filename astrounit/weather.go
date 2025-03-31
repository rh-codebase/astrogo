// weather related functions
package astrounit

import (
	"math"
)

type ScaleHeight struct {
	/**
	 * Infrared scale height
	 * Units: m
	 */
	Ir float64

	/**
	 * dry scale height
	 * Units: m
	 */
	Dry float64

	/**
	 * wet scale height
	 * Units: m
	 */
	Wet float64
}

var (
	scaleHeight  ScaleHeight
	MinElevation Angle
)

type DewPointMethod int

const (
	DP_None DewPointMethod = iota
	CLAUSIUS_CLAPEYRON
	MAGNUS_TETENS
	HOFFMAN_WELCH
)

const (
	/**
	 * Coefficient for dry air pressure in Smith-Weintraub
	 * equation. See, e.g. TMS chapter 13., eqn 13.70
	 * This is K1/Zd with Zd=1. Zd is the dry air compressibility
	 * factor.
	 */
	SW_DRY_AIR = 77.6

	/**
	 * Coefficient for dry air pressure in Smith-Weintraub
	 * equation as measured by Brussaard & Watson (1995)
	 */
	BW_DRY_AIR = 77.6

	/**
	 * Coefficient for water vapor pressure in Smith-Weintraub
	 * equation. Accounts for molecules with induced dipole moments.
	 * See, e.g. TMS chapter 13., eqn 13.70
	 * This is (K2/Zv)/100 with Zv=1. Zv is the H2O gas
	 * compressibility factor.  Note we divide by 100 because we
	 * use relative humidity in the calculation of refractivity.
	 * See interferometry package design document equation 14.
	 */
	SW_INDUCED_DIPOLE = 0.128
	/**
	 * as measured by Brussaard & Watson (1995)
	 */
	BW_INDUCED_DIPOLE = 5.6

	/**
	 * Coefficient for water vapor pressure in Smith-Weintraub
	 * equation. Accounts for molecules with permanent dipole
	 * moments.  See, e.g. TMS chapter 13., eqn 13.70
	 * This is (K3/Zv)/100 with Zv=1. Zv is the H2O gas
	 * compressibility factor.  Note we divide by 100 because
	 * we use relative humidity in the calculation of refractivity.
	 * See interferometry package design document equation 14.
	 */
	SW_PERM_DIPOLE = 3.776e3
	/**
	 * as measured by Brussaard & Watson (1995)
	 */
	BW_PERM_DIPOLE = 3.75e3

	/**
	 * Coefficient for first-order term of frequency
	 * dependence of refractivity, as derived in
	 * interferometry package
	 * design document (see equation 22 and Table 1)
	 */
	REFRAC_COEFF_A = -2.8354e-5

	/**
	 * Coefficient for second-order term of frequency
	 * dependence of refractivity, as derived in
	 * interferometry package
	 * design document (see equation 22 and Table 1)
	 */
	REFRAC_COEFF_B = 5.4012e-7

	/**
	 * Minimum frequency to apply the
	 * frequency dependence of the refractivity.
	 * At or below this frequency the multiplicative
	 * factor is unity.  This value is in GHz because
	 * the equation coefficients are for GHz.
	 */
	MIN_REFRAC_FREQ = 50.0 //50 GHz

	/**
	 * Refraction coefficients from Yan equations 13 and 17.
	 * These are used in the calculation of the refraction
	 * pointing correction. There are 4 equations, 2 for radio
	 * refraction and 2 for optical refraction. Each of those
	 * equations has several coefficients.  The naming scheme
	 * for these coefficients is
	 *
	 * A1_OPT_XX...etc, where <br>
	 *  A1 means first "A1" expression<br>
	 *  RAD means radio <br>
	 *  OPT means optical<br>
	 *  XX indicates the quantity to which
	 *  this coefficient applies:<br>
	 *   "1"   means unity (i.e. its a constant)
	 *   "P"   means atmospheric pressure (measured at ground)<br>
	 *   "T"   means atmospheric temperature (measured at ground)<br>
	 *   "PPW" means the partial pressure of water vapor<br>
	 *   "WAVE" means wavelength<br>
	 *   "SQ" means quantity squared<br>
	 */
	// Equations 13. (radio)
	A1_RAD_1     = 0.5753868  // constant
	A1_RAD_P     = 0.5291e-4  // * P
	A1_RAD_PPW   = -0.2819e-4 // * Pwater
	A1_RAD_PPWSQ = -0.9381e-6 // * Pwater^2
	A1_RAD_T     = -0.5958e-3 // * T
	A1_RAD_TSQ   = 0.2657e-5  // * T^2
	A2_RAD_1     = 1.301211   // constant
	A2_RAD_P     = 0.2003e-4  // * P
	A2_RAD_PPW   = -0.7285e-4 // * Pwater
	A2_RAD_PPWSQ = 0.2579e-5  // * Pwater^2
	A2_RAD_T     = -0.2595e-2 // * T
	A2_RAD_TSQ   = 0.8509e-5  // * T^2

	// Equations 17. (optical)
	A1_OPT_1      = 0.5787089  // constant
	A1_OPT_P      = 0.5609e-4  // * P
	A1_OPT_T      = -0.6229e-3 // * T
	A1_OPT_TSQ    = 0.2824e-5  // * T^2
	A1_OPT_PPW    = 0.5177e-3  // * Pwater
	A1_OPT_PPWSQ  = 0.2900e-6  // * Pwater^2
	A1_OPT_WAVE   = -0.1644e-1 // * lambda
	A1_OPT_WAVESQ = 0.4910e-1  // * lambda^2

	A2_OPT_1      = 1.302474   // constant
	A2_OPT_P      = 0.2142e-4  // * P
	A2_OPT_T      = 0.1287e-2  // * Pwater
	A2_OPT_TSQ    = 0.6500e-6  // * Pwater^2
	A2_OPT_PPW    = -0.2623e-2 // * T
	A2_OPT_PPWSQ  = 0.8776e-5  // * T^2
	A2_OPT_WAVE   = -0.6298e-2 // * lambda
	A2_OPT_WAVESQ = 0.1890e-1  // * lambda^2

	// Zero point offsets used by Mangum (2001)
	// Same as Yan (1996) except temperature has been converted
	// to Kelvin. These are all to be ADDED to the measured
	// values.
	//
	// pressure offset, mbar
	P_OFFSET = -1013.25
	// temperature offset, K
	T_OFFSET = -258.15
	// wavelength offset, micron
	WAVE_OFFSET = -0.5320
	/**
	 * The minimum elevation at which pathlength will be directly
	 * calculated, in radians.  Below this elevation,
	 * computeHorizonPathLength() is used. Currently set to
	 * 1 degree.
	 */
	MIN_ELEVATION = 0.0174533 // radians

	// coefficients for normalized effective zenith
	// Yan equation 12, Mangum equations 24 and 27.
	// R = 8314.14 J/(kmol K) = kg/(m^2 s^2 kmol K)
	// M = 28.970 kg/kmol
	// g = 9.7840 m/s^2 (in vertical column of air at tropopause)
	// R/(M*g) = 29.33895 meters
	// From units:
	// You have: 1 K*gasconstant/(28.970 kg kmol^-1*9.7840 m s^-2)
	// You want:
	//         Definition: 29.333895 m
	RMG = 29.333895
)

func init() {
	scaleHeight.Ir = 2216.0
	scaleHeight.Dry = 9480.0
	scaleHeight.Wet = 2103.0
	MinElevation = NewAngle(Radian, MIN_ELEVATION)
}

/**
 * Compute the refraction pointing correction, given
 * the <em>in vacuo</em> elevation and weather
 * parameters. The return value should be <b>added</b>
 * to the <em>in vacuo</em> elevation.
 * The frequency parameter indicates whether
 * the optical or radio correction is returned. If frequency
 * is greater than 3 THz, the optical correction is returned,
 * otherwise the radio correction is returned.  The refraction
 * correction is done using the mapping function of Yan (1996)
 * as written by Mangum ALMA Memo 366.
 *
 * @param airTemp The ambient air temperature in Kelvin
 * @param atmPressure The atmospheric pressure in millibars
 * @param relHumid The relative humidity in percent
 * @param elevation The <em>in vacuo</em>
 *        (uncorrected) elevation in radians
 * @param frequency The observing frequency, Hz
 * @param  the altitude (station height) of the antenna above
 *         mean Earth radius
 *
 * @return the refraction pointing correction in radians
 * to add to the uncorrected elevation.
 */
// TODO: finish port fix-up WIP
func ComputeRefractionCorrection(airTemp Temperature,
	atmPressure Pressure,
	relHumid float64,
	elevation Angle,
	frequency float64,
	altitude Length) (Angle, error) {
	// Until I figure out the problem with reproducing Mangum's figure
	// use the standard NRAO correction, which I have verified.

	return ulichRefractionCorrection(airTemp, atmPressure,
		relHumid, elevation, frequency)
	/*
		    double zenithRefrac = computeZenithRefractivity(airTemp,atmPressure,
			                                       relHumid, frequency)
		    double mapf  = computeMappingFunction(airTemp, atmPressure,
			                                  relHumid, elevation, frequency,
							  altitude)

		    // return value in radians which should be ADDED to the in vacuo
		    // elevation.  10^-6 scaling is to go from refractivity to refraction
		    // angle.
		    double rcorr = zenithRefrac * cos(elevation) * mapf * micro();
		    //cout << "#YANG MAPF*COS(E): "<< mapf*cos(elevation) << endl;

		    return rcorr;
		    }
	*/
}

// RADIO ONLY: Test against Ulich(1981) radio refraction formulation.
// relHumid in percent. frequency in Hz.
func ulichRefractionCorrection(airTemp Temperature,
	atmPressure Pressure,
	relHumid float64,
	elevation Angle,
	frequency float64) (Angle, error) {

	var ur Angle
	zr, err := ZenithRefractivity(airTemp, atmPressure, relHumid, frequency)
	if err != nil {
		return ur, err
	}
	mapf := ulichMappingFunction(elevation)
	//cout << "#ULICH MAPF: "<< mapf << endl;

	rcorr := zr * mapf * MicroF
	ur = NewAngle(Radian, rcorr)
	return ur, nil

}

// Returns normalized(unitless) effective zenith parameter.
func ulichMappingFunction(elev Angle) float64 {
	sinE := elev.Sin()
	cosE := elev.Cos()
	tanE := math.Tan((87.5 * math.Pi / 180.0) - elev.Radian().Value)
	return cosE / (sinE + 0.00175*tanE)
}

/**
 * The refractivity, N_O, at zenith. See equation 14 of
 * the interferometry design document. Its value depends on
 * weather conditions. This method will NOT check that
 * the air temperature, atmospheric pressure,  and RH
 * are "safe"; use safeAirTemperature(), safeAtmPressure(),
 * and safeRelativeHumidity() to do that.
 *
 * @param airTemp The ambient air temperature in Kelvin
 * @param atmPressure The atmospheric pressure in millibars
 * @param relHumid The relative humidity in percent
 * @param frequency The observing frequency, Hz
 * @return zenith refractivity.
 *
 * @see #safeAirTemperature(double airTemp)
 * @see #safeAtmPressure(double airTemp)
 * @see #safeRelativeHumidity(double relHumid)
 */
func ZenithRefractivity(airTemp Temperature,
	atmPressure Pressure,
	relHumid_pct float64,
	frequency_hz float64) (float64, error) {

	airTK := airTemp.ToKelvin().Value
	var zenithRefrac float64
	// partial pressure of water vapor.
	// NB: the division by 100 to change percent humidity to a number
	// between 0 and 1 is already taken care of in the
	// fixed SMITH_WEINTRAUB coefficients.
	sp, err := WaterSaturatedPressure(airTemp)
	if err != nil {
		return 0.0, err
	}
	ppw := relHumid_pct * sp.ToMillibar().Value

	if IsOptical(frequency_hz) {
		//For now, we do the simple BIMA/OVRO optical refraction
		//which is the Smith-Weintraub equation w/o the permanent
		//dipole term.  Future enhancement will use more sophisticated
		//formulation.
		zenithRefrac = SW_DRY_AIR*atmPressure.ToMillibar().Value/airTK -
			SW_INDUCED_DIPOLE*ppw/airTK
	} else {
		// equation 14 from design doc. (Smith Weintraub equation)
		zenithRefrac = SW_DRY_AIR*atmPressure.ToMillibar().Value/airTK -
			SW_INDUCED_DIPOLE*ppw/airTK +
			SW_PERM_DIPOLE*ppw/(airTK*airTK)
	}

	// multiply by frequency dependence equation
	// NB: We have decided to postpone this correction until
	// after first light, maybe indefinitely.
	//zenithRefrac *= freqDependence(frequency);

	return zenithRefrac, nil
}

// SaterSaturatedPressure returns the saturated air pressure
// using the Clausius-Clapeyron equation.
// See TMS pp512 eq. 13.15
// Bad input and output is checked. NaN, +/-Inf
func WaterSaturatedPressure(airtemp Temperature) (Pressure, error) {
	var pv Pressure

	// check input. No NaN, +/-Inf
	err := SanityCheckF(airtemp.Value, "Illegal Input Temperature")
	if err != nil {
		return pv, err
	}
	airTempK := airtemp.ToKelvin().Value
	tempRatio := (airTempK + AbsoluteZeroCelsius) / airTempK
	scaledTemp := -airTempK / AbsoluteZeroCelsius

	pv.Value = CC_A * math.Exp(CC_B*tempRatio-CC_C*math.Log(scaledTemp))
	pv.Unit = Millibar

	// check output. No NaN, +/-Inf
	err = SanityCheckF(pv.Value, "Illegal output Pressure")
	if err != nil {
		return pv, err
	}
	return pv, nil
}

// WaterPartialPressure returns partial pressure of water vapor.
// relHumid in percent.
func WaterPartialPressure(airTemp Temperature, relHumid float64) (Pressure, error) {
	var pr Pressure
	// divide by 100 to go from RH% -> (0..1)
	wsp, err := WaterSaturatedPressure(airTemp)
	if err != nil {
		return pr, err
	}

	pr.Value = (relHumid / 100.0) * wsp.ToPascal().Value
	pr.Unit = Millibar
	return pr, nil
}

// WaterVaporDensity returns water vapor density, in g/m^3
// TODO: Create Mass, Volume and Density types.
func WaterVaporDensity(airTemp Temperature, relHumid float64) (float64, error) {
	// We can calcuated the density of water vapor from the ideal
	// gas law:
	// pV = m R_i T
	// where R_i is the individual gas constant for the molecule in question.
	// p is partial pressure of water
	// R_i (water) = 461.5 J/kg*K
	//
	// density = m/V, so
	// density = p/(R_i T)

	airTempK := airTemp.ToKelvin().Value
	ppw, err := WaterPartialPressure(airTemp, relHumid)
	if err != nil {
		return 0.0, err
	}

	// Since R_WATER is in SI, we convert pressure to Pascals, which will give
	// the density in kg/m^3.  We want final answer in g/m^3, so two
	// conversions are needed.  Yes, I know the last one is exactly 1000, but
	// code readabilty trumps all.  Besides, we may change units later.
	//factor := units_convert("mbar","Pa") * units_convert("kg/m^3","g/m^3");
	factor := KiloF // kg to gr
	return (factor * ppw.ToPascal().Value / (R_Water * airTempK)), nil
}

// WaterColumn returns total water column, in mm
func WaterColumn(airTemp Temperature, relHumid float64) (Length, error) {
	// Yes, we divide by 1 here, which is no-op, but it is for
	// the sake of code readability.
	var wc Length
	wvd, err := WaterVaporDensity(airTemp, relHumid)
	if err != nil {
		return wc, err
	}
	value := wvd * scaleHeight.Wet / Rho_Water

	// conversion to millimeters from our current funny units.
	//double factor = convert("m (g/m^3)/(g/cm^3)","mm");
	factor := MilliF
	wc.Value = value * factor
	wc.Unit = Millimeter
	return wc, nil

}

func DewPoint(airTemp Temperature, relHumid float64,
	method DewPointMethod) (Temperature, error) {

	var dewTemp Temperature
	dewTemp.Unit = Kelvin

	switch method {
	case CLAUSIUS_CLAPEYRON:
		// partial pressure of water vapor.
		ppw, err := WaterPartialPressure(airTemp, relHumid)
		if err != nil {
			return dewTemp, err
		}
		b := AWIPS_C15 - math.Log(ppw.ToMillibar().Value)
		dewTemp.Value = (b - math.Sqrt(b*b-AWIPS_C3)) / AWIPS_C4

	case MAGNUS_TETENS:
		// The Magnus-Tetens formula is explained here:
		// http://www.paroscientific.com/dewpoint.htm
		// Good over the range Tc = 0 to 60 C, Tdew = 0 to 50 C.
		// Above no longer available. Now using:
		// Lawrence, Mark G. (1 February 2005). "The Relationship between Relative Humidity and the Dewpoint Temperature in Moist Air: A Simple Conversion and Applications". Bulletin of the American Meteorological Society. 86 (2): 225–234. Bibcode:2005BAMS...86..225L. doi:10.1175/BAMS-86-2-225. Retrieved 15 March 2024.

		tc := airTemp.ToCelsius().Value
		alpha := MT_B*tc/(MT_C+tc) + math.Log(relHumid/100.0)
		dewTemp, err := NewTemperature(Celsius, MT_C*alpha/(MT_B-alpha)-AbsoluteZeroCelsius)
		if err != nil {
			return dewTemp, err
		}

	case HOFFMAN_WELCH:
		// Hoffman & Welch from HatCreek's weatherman1.c code.
		// This is the integrated form Clausius Clapeyron equation where
		// one temperature bound is set to infinity.
		// ln(P1/P2) = H/R (1/T2 - 1/T1)
		// let T2 = infinity, P2 = P0
		//
		// P = P0*exp(-H/RT)
		// where P is saturated water vapor pressure
		//       P0 is vapor pressure of water at infinite temperature
		//       H is the enthalpy of evaporation of water
		//       R is the gas constant

		// pp of water, [mm Hg]
		ppw := NewPressure(MillimeterHg, HW_A*math.Exp(-HW_B/airTemp.ToCelsius().Value)*(relHumid/100.0))
		dewTemp.Value = -HW_B / math.Log(ppw.ToMillibar().Value/HW_A) // Kelvin

	}

	return dewTemp, nil
}

func RelativeHumidity(airTemp, dewTemp Temperature, method DewPointMethod) (float64, error) {

	var relHumid float64
	Tc := airTemp.ToCelsius().Value
	Td := dewTemp.ToCelsius().Value

	switch method {
	default:
	case CLAUSIUS_CLAPEYRON:
		// This is the definition of relative humidity!
		dtwsp, err := WaterSaturatedPressure(dewTemp)
		if err != nil {
			return relHumid, err
		}
		atwsp, err := WaterSaturatedPressure(airTemp)
		if err != nil {
			return relHumid, err
		}
		relHumid = 100.0 * dtwsp.ToMillibar().Value / atwsp.ToMillibar().Value

	case MAGNUS_TETENS:
		ratioA := (MT_B * Tc) / (MT_C + Tc)
		ratioB := (MT_B*Td - Td*ratioA - MT_C*ratioA) / (MT_C + Td)
		relHumid = 100.0 * math.Exp(ratioB)

	case HOFFMAN_WELCH:
		// Hoffman & Welch from HatCreek's weatherman1.c code
		Ph2o := HW_A * math.Exp(-HW_B/Td)
		relHumid = 100. * Ph2o * math.Exp(HW_B/Tc) / HW_A

	}
	return relHumid, nil
}

// PathLength returns refractivity integrated through the atmosphere,
// that is, the pathlength. altitude is the antenna height above
// mean Earth Radius
func Pathlength(airTemp Temperature,
	atmPressure Pressure,
	relHumid float64,
	elevation Angle,
	frequency float64,
	altitude Length) (Length, error) {

	var pl Length
	// If we are observing "over the top", then the
	// elevation can be larger than 90 degrees.
	// This will give a negative tan(elev) value and
	// therefore a negative pathlength, which is unphysical.
	// So put elevation into first quadrant if needed.
	// This is okay since the atmosphere is assumed to
	// be azimuthally symmetric.
	if elevation.Degree().Value > 90.0 {
		elevation.Value = 180. - elevation.Degree().Value
		elevation.Unit = Degree
	}

	// If we are below the minimum elevation,
	// then return L_max as defined by TMS eqn 13.40
	// At 7 degrees, we underpredict L by 2.2% for a standard atmosphere.
	if elevation.LessThan(MinElevation) {
		//cout << "using horizon path length" << endl;
		return HorizonPathlength(airTemp, atmPressure,
			relHumid, frequency)
	}

	// leading coefficient, 1/sin(E) safe because
	// we've already eliminated low elevations
	factor := MicroF / (airTemp.ToCelsius().Value * elevation.Sin())

	// Cotangent squared of elevation.
	// Do it by cos/sin instead of 1/tan because tan
	// diverges at 90 degrees.  cos/sin is safe because
	// we have already eliminated small elevations where
	// sin(E) would diverge.
	cotE := elevation.Cos() / elevation.Sin()
	cotsqE := cotE * cotE

	// saturated H2O pressure, millibars
	satPressure, err := WaterSaturatedPressure(airTemp)
	if err != nil {
		return pl, err
	}

	// partial pressure times 100.
	// NB: the division by 100 to change percent humidity to a number
	// between 0 and 1 is already taken care of in the
	// fixed SMITH_WEINTRAUB coefficients.
	ppw := relHumid * satPressure.ToMillibar().Value

	// add the antenna height above the ground plane to
	// the earth radius to get total height above the
	// center of the earth.
	adjustedRadius := EARTH_RADIUS + altitude.Meter().Value

	// compute three terms in the equation separately.
	// Since scaleheight and Earth radius are in meters
	// the resultant pathlength will be in meters.
	term1 := SW_DRY_AIR * atmPressure.ToMillibar().Value * scaleHeight.Dry
	term1 *= (1.0 - cotsqE*scaleHeight.Dry/adjustedRadius)

	// ----------------- TEST: "EXACT" FROM TMS -----------
	// Attempt to reproduce TMS figure 13.35, making a guess as
	// to what they used for atmPressure, scaleHeight.dry, and EARTH.radius.
	// This should give about L=2.31 meters at zenith.
	//term1 = SW_DRY_AIR * DEFAULT_ATM_PRESSURE * 8.0;
	//term1 *= (1.0 - cotsqE * 8.0 / 6370.0);
	// ---------------- TEST--------------------------

	term2 := -SW_INDUCED_DIPOLE * ppw * scaleHeight.Wet
	term2 *= (1.0 - cotsqE*scaleHeight.Wet/adjustedRadius)

	term3 := SW_PERM_DIPOLE * ppw * scaleHeight.Ir
	term3 *= (1.0 - cotsqE*scaleHeight.Ir/adjustedRadius)
	term3 /= airTemp.ToCelsius().Value

	// -------------- TEST EXACT FROM TMS --------------------------
	// In figure 13.35, they set rho_v to zero, which is equivalent
	// to ignoring the second and third terms.
	//term2 = 0.;
	//term3 = 0.;
	// -------------- TEST EXACT FROM TMS --------------------------

	// pathlength in meters
	pl.Value = factor * (term1 + term2 + term3)
	pl.Unit = Meter

	// multiply by frequency dependence equation.
	// NB: We have decided to postpone this correction until
	// after first light
	//pathlength *= freqDependence(frequency);

	return pl, nil
}

func HorizonPathlength(airTemp Temperature, atmPressure Pressure,
	relHumid float64, frequency float64) (Length, error) {
	var pl Length
	// First, get the zenith refractivity.
	pathlength, err := ZenithRefractivity(airTemp, atmPressure,
		relHumid,
		frequency)
	if err != nil {
		return pl, err
	}

	//double zrefrac = pathlength;

	// TMS equation 13.40 computes horizon pathlength in
	// whatever units Earth radius and atmospheric scale height
	// happen to be in, and uses a coefficient out front of 10^-6.
	// Note we use the "wet" scale height, since TMS uses 2 kilometers.

	factor := math.Pi * EARTH_RADIUS * scaleHeight.Wet / 2.0
	pathlength *= MicroF * math.Sqrt(factor)

	// ratio should be about 0.14 for pathlength in meters,
	// before applying frequency dependence.
	//cout << " pathlength/zrefrac = " << pathlength/zrefrac;

	// TODP: multiply by frequency dependence equation.
	//pathlength *= freqDependence(frequency);

	pl.Value = pathlength
	pl.Unit = Meter
	return pl, nil
}

// IsOptical return true if frequency greater than 1THz. Frequency specified
// in Hertz.
func IsOptical(frequency_hz float64) bool {
	if frequency_hz > 3.0*TeraF {
		return true
	} else {
		return false
	}
}

/*
double waterPartialPressure(double airTemp, double relHumid)
{
    // partial pressure of water vapor.
    // divide by 100 to go from RH% -> (0..1)
    return ( relHumid / 100.0)  * computeSaturatedPressure(airTemp);
}

// returns total water column, in mm
double waterColumn( double airTemp, double relHumid)
{
  // Yes, we divide by 1 here, which is no-op, but it is for
  // the sake of code readability.
  const double rho_water = 1.0;   // density of water in g/cm^3
  double value = waterVaporDensity(airTemp, relHumid)
                 * scaleHeight.wet / rho_water;

  // conversion to millimeters from our current funny units.
  //double factor = convert("m (g/m^3)/(g/cm^3)","mm");
  double factor = 0.001;
  return ( value  * factor );

}

double computeDewPoint(double airTemp, double relHumid,
	                           dewPointMethodType method)
{
  double dewTemp;
  double alpha, b, ppw, Tc;

  switch (method) {

      default:
      case CLAUSIUS_CLAPEYRON:
	// partial pressure of water vapor.
	ppw = waterPartialPressure(airTemp, relHumid);
	b = AWIPS_C15 - log(ppw);
	dewTemp = ( b - sqrt(b*b-AWIPS_C3) )/AWIPS_C4;
	break;

      case MAGNUS_TETENS:
    // The Magnus-Tetens formula is explained here:
    // http://www.paroscientific.com/dewpoint.htm
    // Good over the range Tc = 0 to 60 C, Tdew = 0 to 50 C.
	Tc = airTemp + ABS_ZERO;
	alpha = MT_A * Tc/(MT_B + Tc) + log(relHumid/100.0);
	dewTemp = MT_B *alpha/(MT_A - alpha) - ABS_ZERO;
	break;

      case HOFFMAN_WELCH:
	// Hoffman & Welch from HatCreek's weatherman1.c code.
	// This is the integrated form Clausius Clapeyron equation where
	// one temperature bound is set to infinity.
	// ln(P1/P2) = H/R (1/T2 - 1/T1)
	// let T2 = infinity, P2 = P0
	//
	// P = P0*exp(-H/RT)
	// where P is saturated water vapor pressure
	//       P0 is vapor pressure of water at infinite temperature
	//       H is the enthalpy of evaporation of water
	//       R is the gas constant

	ppw = HW_A * exp(-HW_B/airTemp) * (relHumid/100.0);  // pp of water, [mm Hg]
	dewTemp = -HW_B/log(ppw/HW_A);   // Kelvin
	break;

  }

  return dewTemp;
}


double computeHumidity(double airTemp, double dewTemp,
                       dewPointMethodType method)
{
  double relHumid;
  double Tc = airTemp + ABS_ZERO;
  double Td = dewTemp + ABS_ZERO;
  double ratioB, ratioA, Ph2o;
  switch ( method ) {
      default:
      case CLAUSIUS_CLAPEYRON:
        // This is the definition of relative humidity!
	relHumid = 100.0 * computeSaturatedPressure(dewTemp)/computeSaturatedPressure(airTemp);
        break;
      case MAGNUS_TETENS:
	ratioA = (MT_A * Tc)/(MT_B + Tc);
	ratioB = (MT_A * Td - Td * ratioA - MT_B * ratioA)/(MT_B + Td);
	relHumid = 100.0 * exp( ratioB );
        break;
      case HOFFMAN_WELCH:
	// Hoffman & Welch from HatCreek's weatherman1.c code
	Ph2o = HW_A * exp(-HW_B/dewTemp);
	relHumid = 100*Ph2o*exp(HW_B/airTemp)/HW_A;
        break;
  }
  return relHumid;
}

// returns refractivity integrated through the atmosphere,
// that is, the pathlength.
double computePathlength( double airTemp,
				      double atmPressure,
	                              double relHumid,
				      double elevation,
				      double frequency,
				      double altitude
				    )
{

    // If we are observing "over the top", then the
    // elevation can be larger than 90 degrees.
    // This will give a negative tan(elev) value and
    // therefore a negative pathlength, which is unphysical.
    // So put elevation into first quadrant if needed.
    // This is okay since the atmosphere is assumed to
    // be azimuthally symmetric.
    if (elevation > M_PI_2 ) {
	elevation = M_PI - elevation;
    }

    // If we are below the minimum elevation,
    // then return L_max as defined by TMS eqn 13.40
    // At 7 degrees, we underpredict L by 2.2% for a standard atmosphere.
    if (elevation < MIN_ELEVATION) {
	//cout << "using horizon path length" << endl;
	return computeHorizonPathlength(airTemp, atmPressure,
		                        relHumid, frequency);
    }

    // leading coefficient, 1/sin(E) safe because
    // we've already eliminated low elevations
    double factor = micro()/(airTemp*sin(elevation));

    // Cotangent squared of elevation.
    // Do it by cos/sin instead of 1/tan because tan
    // diverges at 90 degrees.  cos/sin is safe because
    // we have already eliminated small elevations where
    // sin(E) would diverge.
    double cotE = cos(elevation)/sin(elevation);
    double cotsqE = cotE*cotE;

    // saturated H2O pressure, millibars
    double satPressure = computeSaturatedPressure(airTemp);

    // partial pressure times 100.
    // NB: the division by 100 to change percent humidity to a number
    // between 0 and 1 is already taken care of in the
    // fixed SMITH_WEINTRAUB coefficients.
    double ppw = relHumid * satPressure;

    // add the antenna height above the ground plane to
    // the earth radius to get total height above the
    // center of the earth.
    double adjustedRadius = EARTH_RADIUS + altitude;

    // compute three terms in the equation separately.
    // Since scaleheight and Earth radius are in meters
    // the resultant pathlength will be in meters.
    double term1 = SW_DRY_AIR * atmPressure * scaleHeight.dry;
    term1 *= (1.0 - cotsqE * scaleHeight.dry / adjustedRadius);

    // ----------------- TEST: "EXACT" FROM TMS -----------
    // Attempt to reproduce TMS figure 13.35, making a guess as
    // to what they used for atmPressure, scaleHeight.dry, and EARTH.radius.
    // This should give about L=2.31 meters at zenith.
    //term1 = SW_DRY_AIR * DEFAULT_ATM_PRESSURE * 8.0;
    //term1 *= (1.0 - cotsqE * 8.0 / 6370.0);
    // ---------------- TEST--------------------------

    double term2 = -SW_INDUCED_DIPOLE * ppw * scaleHeight.wet;
    term2 *= (1.0 - cotsqE * scaleHeight.wet / adjustedRadius);

    double term3 = SW_PERM_DIPOLE * ppw * scaleHeight.ir;
    term3 *= (1.0 - cotsqE * scaleHeight.ir / adjustedRadius);
    term3 /= airTemp;

    // -------------- TEST EXACT FROM TMS --------------------------
    // In figure 13.35, they set rho_v to zero, which is equivalent
    // to ignoring the second and third terms.
    //term2 = 0.;
    //term3 = 0.;
    // -------------- TEST EXACT FROM TMS --------------------------

    // pathlength in meters
    double pathlength = factor*(term1 + term2 + term3);

    // multiply by frequency dependence equation.
    // NB: We have decided to postpone this correction until
    // after first light
    //pathlength *= freqDependence(frequency);

    //cout << "elev " << elevation << endl;
    //cout << "factor " << factor << endl;
    //cout << "cotsqE " << cotsqE << endl;
    //cout << "satPres " << satPressure << endl;
    //cout << "term1 " << term1 << endl;
    //cout << "term2 " << term2 << endl;
    //cout << "term3 " << term3 << endl;
    //cout << "freqDep " << freqDependence(frequency) << endl;
    //cout << "pathlength (m) " << pathlength << endl;


    return pathlength ;
}

double computeHorizonPathlength( double airTemp,
	                                     double atmPressure,
					     double relHumid,
					     double frequency)
{
    // First, get the zenith refractivity.
    double pathlength = computeZenithRefractivity(airTemp,
	                                          atmPressure,
	 					  relHumid,
						  frequency);

    //double zrefrac = pathlength;

    // TMS equation 13.40 computes horizon pathlength in
    // whatever units Earth radius and atmospheric scale height
    // happen to be in, and uses a coefficient out front of 10^-6.
    // Note we use the "wet" scale height, since TMS uses 2 kilometers.

    double factor = M_PI * EARTH_RADIUS * scaleHeight.wet / 2.0;
    pathlength *= sqrt(factor);
    pathlength *= micro();

    // ratio should be about 0.14 for pathlength in meters,
    // before applying frequency dependence.
    //cout << " pathlength/zrefrac = " << pathlength/zrefrac;

    // multiply by frequency dependence equation.
    // NB: We have decided to postpone this correction until
    // after first light
    //pathlength *= freqDependence(frequency);

    return pathlength;

}



//------------------------------------------------------
//            PRIVATE METHODS
//------------------------------------------------------
double freqDependence( double freq )
{
    // convert Hz to GHz
    freq *= nano();

    // If freq is less than to fiducial frequency,
    // then the scale factor is unity.
    //std::cout.setf(ios::fixed);
    //std::cout.precision(8);
    if ( freq < MIN_REFRAC_FREQ) {
	//std::cout << "Freq " << freq << " <= " << MIN_REFRAC_FREQ << endl;

	return 1.0;
    } else {
	//std::cout << "Freq " << freq << " > " << MIN_REFRAC_FREQ << endl;
	// scale factor is (1 + A*freq + B*freq^2)
	return ( 1.0 + REFRAC_COEFF_A * freq +
	         REFRAC_COEFF_B * freq * freq
	       );
    }
}

double computeMappingFunction(
	               double airTemp, double atmPressure,
                       double relHumid,
                       double elevation,
		       double frequency,
		       double altitude)
{


    double a1, a2; // the individual refraction expressions.

    // Set up some intermediate values.

    // Be sure to guard against zero elevation
    double sinE = sin(elevation);

    double p0  = atmPressure + P_OFFSET;
    double t0  = airTemp + T_OFFSET;

    // partial pressure of water vapor
    double ppw = waterPartialPressure(airTemp, relHumid);

    double nez = normalizedEffectiveZenith(airTemp, elevation, altitude);

    // If the frequency is greater than 3 THz, use the optical
    // refraction experessions, otherwise use the radio.
    if ( isOptical(frequency) ) {
      double w0 = C/frequency * units_convert("m","micron")
	            + WAVE_OFFSET;
        a1 = a1Optical(p0,t0,ppw,w0);
        a2 = a2Optical(p0,t0,ppw,w0);
    } else {
        a1 = a1Radio(p0,t0,ppw);
        a2 = a2Radio(p0,t0,ppw);
    }


    // Mapping function is a continued fraction. Do it in parts.
    double f     = sinE + MAP1 / (nez + MAP2);
    double value = sinE +   a1 / (nez + a2/f);

   // cout << "#MAP: "
	 // << "temp: "<< t0  << "  "
	 // << "pres: "<< p0  << "  "
	 // << "PPW:" << ppw << "  "
	 // << "freq:" << frequency << "  "
	 // << "wave:" << w0 << "  "
	 // << "Erad: " << elevation << "  "
	 // << "nu: " <<frequency << "  "
	 // << "ALT:" <<altitude  << "  "
	 // << "NEZ: " <<nez    << "  "
   //       << "a1: " <<a1     << "  "
   //       << "a2: " <<a2     << "  "
   //       << "f: " <<f     << "  "
   //       << "return: " <<1.0/value << "  "
	 // << endl;

    return 1.0/value;

}


int isOptical(double frequency) {
    if ( frequency > (3.0 * tera()) ) {
	return 1;
    } else {
	return 0;
    }
}

double normalizedEffectiveZenith(double airTemp,
	                                     double elevation,
					     double altitude)
{

    // the normalized effective zenith is
    // sqrt(r0/2H)*tan(E).
    // However, in all equations it appears as
    // (I^2)*csc(E) = (r0/2H)*tan^2(E)*csc(E)
    //              = (r0/2H)*[sin^2(E)/cos^2(E)]*[1/sin(E)]
    //              = (r0/2H)*sin(E)/cos^2(E).
    //
    //This quantity diverges at cos(E)=0, or E=90 degrees.

    // Check for elevation = 90, which would cause divergence.
    // If so, just return a large number.
    double dLimits_epsilon = .01;
    double dLimits_max = 1.e99;
    double diff = abs(elevation - M_PI/2);
    if(diff <= dLimits_epsilon) {
	// Return a number as close to infinity as we are allowed.
	// returning a large number here is fine because
	// it will be put in the denominators of the
	// continued fraction mapping function. And besides
	// the whole thing gets multipied by cos(elevation)
	// in the end, which is zero.
	return dLimits_max;
    }

    double effHeight  = airTemp * RMG; // effective height, in meters
    double fact = (altitude + EARTH_RADIUS)/(2.0*effHeight);

    // cout << "#In nez: Erad: "<< elevation
	// << " airTemp:   " << airTemp
	// << " radius:   " << altitude + EARTH_RADIUS
	// << " effheight: " << effHeight
	// << " fact: " << fact << endl;

    return fact * sin(elevation)/pow(cos(elevation),2.0);
}

double a1Radio(double p0, double t0, double ppw)
{
    double value = A1_RAD_1 + A1_RAD_P*p0 + A1_RAD_PPW*ppw
                 + A1_RAD_PPWSQ*ppw*ppw + A1_RAD_T*t0
                 + A1_RAD_TSQ*t0*t0;
    return value;
}

double a2Radio(double p0, double t0, double ppw)
{
    double value = A2_RAD_1 + A2_RAD_P*p0 + A2_RAD_PPW*ppw
                 + A2_RAD_PPWSQ*ppw*ppw + A2_RAD_T*t0
                 + A2_RAD_TSQ*t0*t0;
    return value;

}

double a1Optical(double p0, double t0, double ppw, double w0)
{
    double value = A1_OPT_1 + A1_OPT_P*p0 + A1_OPT_PPW*ppw
                 + A1_OPT_PPWSQ*ppw*ppw + A1_OPT_T*t0
                 + A1_OPT_TSQ*t0*t0 + A1_OPT_WAVE*w0
                 + A1_OPT_WAVESQ*w0*w0;
    return value;

}

double a2Optical(double p0, double t0, double ppw, double w0)
{
    double value = A2_OPT_1 + A1_OPT_P*p0 + A2_OPT_PPW*ppw
                 + A2_OPT_PPWSQ*ppw*ppw + A2_OPT_T*t0
                 + A2_OPT_TSQ*t0*t0 + A2_OPT_WAVE*w0
                 + A2_OPT_WAVESQ*w0*w0;
    return value;

}

double micro() {
  return 1.e-6;
}

double nano() {
  return 1.e-9;
}

double tera() {
  return 1.e12;
}

double units_convert(const char* u1, const char* u2) {

  // convertable units

  // mbar conversion
  if (strncmp(u1, "mbar", 4) == 0) {
    if (strncmp(u2, "mbar", 4) == 0) {
      return 1.0;
    } else if (strncmp(u2, "Pa", 2) == 0) {
      return 100.0;
    }

    // Pa convertion
  } else if (strncmp(u1, "Pa", 2) == 0) {
    if (strncmp(u2, "mbar", 4) == 0) {
      return .01;
    } else if (strncmp(u2, "Pa", 2) == 0) {
      return 1.0;
    }

    // kg/m^3 convertion
  } else if (strncmp(u1, "kg/m^3", 6) == 0) {
    if (strncmp(u2, "kg/m^3", 6) == 0) {
      return 1.0;
    } else if (strncmp(u2, "g/m^3", 5) == 0) {
      return 1000.0;
    }

    // g/m^3 conversion
  } else if (strncmp(u1, "g/m^3", 5) == 0) {
    if (strncmp(u2, "g/m^3", 5) == 0) {
      return 1.0;
    } else if (strncmp(u2, "kg/m^3", 6) == 0) {
      return 0.001;
    }

    // m conversion
  } else if (strncmp(u1, "m", 1) == 0) {
    if (strncmp(u2, "m", 1) == 0) {
      return 1.0;
    } else if (strncmp(u2, "cm", 2) == 0) {
      return 100;
    } else if (strncmp(u2, "mm", 2) == 0) {
      return 1000;
    } else if (strncmp(u2, "micron", 6) == 0) {
      return 1e6;
    }

    // micron conversion
  } else if (strncmp(u1, "micron", 6) == 0) {
    if (strncmp(u2, "micron", 6) == 0) {
      return 1.0;
    } else if (strncmp(u2, "m", 1) == 0) {
      return 1.0e-6;
    } else if (strncmp(u2, "cm", 1) == 0) {
      return 1.0e-4;
    } else if (strncmp(u2, "mm", 1) == 0) {
      return 1.0e-3;
    }
  }
  printf("Can not convert\n");
}

void computeAtmosQuant(float temp, float relHum, atmosQuantType* aqt) {
  float tK = temp - ABS_ZERO; // convert Centigrade to Kelvin
  aqt->wvd = waterVaporDensity(tK, relHum);
  aqt->pw  = waterColumn(tK, relHum);
  aqt->dewPoint = computeDewPoint(tK, relHum, CLAUSIUS_CLAPEYRON) + ABS_ZERO;
}

void computeMinMax(atmosQuantType current,
                   atmosQuantType* lo,
                   atmosQuantType* hi) {
  // Water Vapor Denity
  if (current.wvd < lo->wvd) {
    lo->wvd = current.wvd;
    setTime(&(lo->wvdTime));
  }
  if (current.wvd > hi->wvd) {
    hi->wvd = current.wvd;
    setTime(&(hi->wvdTime));
  }

  // Precipital Water
  if (current.pw < lo->pw) {
    lo->pw = current.pw;
    setTime(&(lo->pwTime));
  }
  if (current.pw > hi->pw) {
    hi->pw = current.pw;
    setTime(&(hi->pwTime));
  }

  // Dew Point
  if (current.dewPoint < lo->dewPoint) {
    lo->dewPoint = current.dewPoint;
    setTime(&(lo->dewPointTime));
  }
  if (current.dewPoint > hi->dewPoint) {
    hi->dewPoint = current.dewPoint;
    setTime(&(hi->dewPointTime));
  }
}

void setTime(WX_SENS* tt) {
  time_t secs = time(0);
  struct tm* t;

  t = localtime(&secs);
  tt->mon  = t->tm_mon + 1;
  tt->day  = t->tm_mday;
  tt->hour = t->tm_hour;
  tt->min  = t->tm_min;
  tt->sec  = t->tm_sec;
}
*/
