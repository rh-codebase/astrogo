// math functions
package astrounit

import (
	"errors"
	"fmt"
	"math"
	"time"

	at "github.com/rh-codebase/astrogo/astrotime"
	"gonum.org/v1/gonum/mat"

	"github.com/montanaflynn/stats"
)

const (
	//	voltsPerDeg = 0.0011
	voltsPerDeg = 0.06467 / 360.0
	// wrap voltage at az=0, wrap 0
	voltZero = 2.73
)

var (
	wrapZero Angle = NewAngle(Degree, 0.0)
)

// Solve m*x = b for x. x = mInv * b
// TODO: move to astrounit/math.go as this is general
func Solve(m [][]float64, b []float64) *mat.VecDense {
	l := len(b)
	dm := make([]float64, l*l)
	kdx := 0
	for rdx := 0; rdx < l; rdx++ {
		for cdx := 0; cdx < l; cdx++ {
			dm[kdx] = m[rdx][cdx]
			kdx++
		}
	}
	dv := mat.NewVecDense(l, b)
	mm := mat.NewDense(l, l, dm)
	mv := mat.NewVecDense(l, nil)
	mv.SolveVec(mm, dv)
	return mv
}

// WrapVoltage outputs the voltage which should display in junction box
// corresponding to the azimuth potentiometer.
func WrapVoltage(az Angle, wc int32) float64 {
	return voltZero + voltsPerDeg*360.*float64(wc) + voltsPerDeg*(az.Degree().Value)
	//return 1.784 + 0.0011*az.Degree().Value
	//return 2.18 + .0011*(az.Degree().Value-360.)
	//return 1.04 + 0.0011*az.Degree().Value
}

// WrapAngle outputs the Wrap Angle
// corresponding to the wrap voltage
// Deprecated.
func WrapAngle(v float64) Angle {
	return NewAngle(Degree, (v-voltZero)/voltsPerDeg)
}

// MapAngle returns the Az angle in wrap space.
func MapAngle(wc int32, ca Angle) Angle {
	a := NewAngle(Degree, 360.*float64(wc))
	return ca.Add(a)
}

// WrapCounterByVolts returns a function which returns the wrap count
// based on the wrap voltage. This is used at startup to determine the
// initial wrap count of the telescope.
func WrapCouterByVolts(voltOffset, voltsPerRev float32) func(float32) int32 {
	vo := voltOffset
	vpr := voltsPerRev
	return func(cv float32) int32 {
		return int32((cv - vo) / vpr)
	}
}

// WrapCount outputs the wrap count given the wrap voltage and the
// current angle. mwc is the maximum wrap count to consider and tol is
// the voltage tolerance +/-
// Deprecated. Use WrapCounterByVolts instead.
func WrapCount(volts float64, voltsPerRev, tol float64, mwc int32) (int32, error) {
	wa := WrapAngle(volts).Degree().Value
	var si int32
	if wa > 0 {
		si = 1
	} else {
		si = -1
	}
	var idx int32
	for idx = 0; idx <= mwc; idx++ {
		ca := float64(idx) * 360.0
		maxA := ca + wrapZero.Degree().Value
		minA := -maxA
		if wa > minA && wa < maxA {
			return int32(idx) * si, nil
		}
	}

	return int32(0), errors.New("Unable to determine wrap count")
}

// Median is a closure which returns the median over the number of samples
// specified.
func MedianInt(samples int) func(int) (int, error) {
	arr := make([]float64, samples)
	var ptr int
	return func(v int) (int, error) {
		arr[ptr] = float64(v)
		ptr += 1
		if ptr >= samples {
			ptr = 0
		}
		m, err := stats.Median(arr)
		return int(m), err
	}
}

// Median is a closure which returns the median over the number of samples
// specified.
func Median(samples int) func(float64) (float64, error) {
	arr := make([]float64, samples)
	var ptr int
	return func(v float64) (float64, error) {
		arr[ptr] = v
		ptr += 1
		if ptr >= samples {
			ptr = 0
		}
		m, err := stats.Median(arr)
		return m, err
	}
}

// Median is a closure which returns the median over the number of samples
// specified.
func MedianAngle(samples int) func(Angle) (Angle, error) {
	arr := make([]float64, samples)
	var ptr int
	return func(v Angle) (Angle, error) {
		arr[ptr] = v.Degree().Value
		ptr += 1
		if ptr >= samples {
			ptr = 0
		}
		m, err := stats.Median(arr)
		return NewAngle(Degree, m), err
	}
}

// WrapCounter returns a function which return the wrap count based on
// the current azimuth angle. It is initialized with the current encoder
// angle and wrap count read at startup.
func WrapCounter(encoderAngle Angle, wc int32) func(Angle) int32 {
	wrapCount := wc
	prevAngle := encoderAngle
	prevAngleT := time.Now()
	return func(encA Angle) int32 {
		fmt.Println("prevA: ", prevAngle.Degree().Value, " encA: ", encA.Degree().Value)
		encAT := time.Now()
		delT := encAT.Sub(prevAngleT)
		fmt.Println("delT: ", delT)
		upperDeg := 360. - delT.Seconds()
		lowerDeg := delT.Seconds()
		prevAv := prevAngle.Degree().Value
		encAv := encA.Degree().Value
		if ((prevAv > upperDeg) && (encAv < lowerDeg)) ||
			((prevAv < lowerDeg) && (encAv > upperDeg)) {
			fmt.Println("angles within acceptable limits")

			delA := encA.Sub(prevAngle)
			delAv := 360. - math.Abs(delA.Degree().Value)
			fmt.Println("delAv: ", delAv)
			rate := NewAngularRate(NewAngle(Degree, delAv), delT)

			if rate.Degree().Value > 2.0 {
				fmt.Println("rate > 2.0. return w/out change")
				return wrapCount
			}

			if encA.GreaterThan(prevAngle) { // moved couterclockwise
				wrapCount--
			} else if encA.LessThan(prevAngle) { // moved clockwise
				wrapCount++
			}
		}
		prevAngle = encA
		prevAngleT = encAT
		return wrapCount
	}
}

// AngleWrapCount is a closure which counts the number of times 2 successive
// angles cross a boundry. +1 count from every boundry-thresh->boundry+thresh
// crossing and -1 count for boundry+thresh -> boundry-thresh crossing.
// Deprecated. Use WrapCounter
func AngleWrapCount(boundry, thresh Angle) func(Angle) int32 {
	prev := float64(-1.0)
	wrapCount := int32(0)
	wrapHighThresh := boundry.SubModulo(thresh, Modulo360).Degree().Value
	wrapLowThresh := boundry.AddModulo(thresh, Modulo360).Degree().Value
	return func(v Angle) int32 {
		curr := v.Degree().Value
		if prev == -1.0 {
			prev = curr
			return 0
		}
		if prev > wrapHighThresh && curr < wrapLowThresh {
			wrapCount++
		} else if prev < wrapLowThresh && curr > wrapHighThresh {
			wrapCount--
		}
		//fmt.Println("doing nothing. prev: ", prev, " curr: ", curr)
		prev = curr
		return wrapCount
	}
}

func SanityCheckF(v float64, msg string) error {
	if math.IsNaN(v) {
		emsg := fmt.Sprintf("%s Value is NaN", msg)
		return errors.New(emsg)
	}
	if math.IsInf(v, 1) {
		emsg := fmt.Sprintf("%s Value is +Inf", msg)
		return errors.New(emsg)
	}
	if math.IsInf(v, -1) {
		emsg := fmt.Sprintf("%s Value is -Inf", msg)
		return errors.New(emsg)
	}
	return nil
}

// Modulo24 returns a number between (0,24]
func Modulo24(hours float64) float64 {
	/*
		if hours > 0.0 && hours < 24.0 {
			return hours
		}
		if hours < 0.0 && hours > -24.0 {
			return hours + 24.0
		}
		rem := int64(math.Mod(hours, 24.0))
		fmt.Println("rem= ", rem)
		return math.Mod(float64(rem)*hours, 24.0)
	*/
	for hours <= at.HourPerDay {
		hours += at.HourPerDay
	}
	for hours >= at.HourPerDay {
		hours -= at.HourPerDay
	}
	return hours
}

// Modulo360 returns a number between (0,360]
func Modulo360(deg float64) float64 {
	for deg <= 360. {
		deg += 360.
	}
	for deg >= 360. {
		deg -= 360.
	}
	return deg
}

// Modulo360 returns a number between (0,360]
func Modulo2Pi(a Angle) Angle {
	v := Modulo360(a.Degree().Value)
	return NewAngle(Degree, v)
}

// ModuloN is a closure returning a function which returns number between (0,n]
func ModuloN(n float64) func(float64) float64 {
	m := func(v float64) float64 {
		for v <= n {
			v += n
		}
		for v >= n {
			v -= n
		}
		return v
	}
	return m
}

// SecondsAfterBoundry computes the current time past the defined boundry.
// boundry represent some relateve time boundry (i.e 0.5 or 0.1) from which
// to compute the current time against. If for example, the current time
// in unix seconds is 1720548692.123 and the boundry is .5, the the
// return value would be 0.123 seconds. If the boundry is 0.1, then 0.023
// seconds is returned.
// ct is the current time, boundry is a Duration like 100*time.Millisecond
func SecondsAfterBoundry(ct time.Time, boundry time.Duration) float64 {
	fct := float64(ct.UnixNano()) * 1e-9
	db := float64(boundry) * 1e-9

	ffct := fct / db
	dt := (ffct - float64(int64(ffct))) * db
	return dt
}

func reverse8Bits(num byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}
