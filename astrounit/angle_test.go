package astrounit

import (
	"fmt"
	"math"
	"testing"

	th "github.com/rh-codebase/genutilsgo"
)

func TestNewAngle(t *testing.T) {
	// Radian Test
	a := NewAngle(Radian, 1.1)
	th.CheckS(t, a.UnitString(), RadianStr, "Unit Error")
	th.CheckF(t, a.Value, 1.1, "Value Error")

	// Hour Test
	a = NewAngle(Hour, 12.1)
	th.CheckS(t, a.UnitString(), HourStr, "Unit Error")
	th.CheckF(t, a.Value, 12.1, "Value Error")

	// Degree Test
	a = NewAngle(Degree, 1.2)
	th.CheckS(t, a.UnitString(), DegreeStr, "Unit Error")
	th.CheckF(t, a.Value, 1.2, "Value Error")

	// MilliRadian Test
	a = NewAngle(MilliRadian, 1.3)
	th.CheckS(t, a.UnitString(), MilliRadianStr, "Unit Error")
	th.CheckF(t, a.Value, 1.3, "Value Error")

	// ArcMinute Test
	a = NewAngle(ArcMinute, 1.4)
	th.CheckS(t, a.UnitString(), ArcMinuteStr, "Unit Error")
	th.CheckF(t, a.Value, 1.4, "Value Error")

	// ArcSecond Test
	a = NewAngle(ArcSecond, 1.5)
	th.CheckS(t, a.UnitString(), ArcSecondStr, "Unit Error")
	th.CheckF(t, a.Value, 1.5, "Value Error")

	// MilliArcSecond Test
	a = NewAngle(MilliArcSecond, 1.6)
	th.CheckS(t, a.UnitString(), MilliArcSecondStr, "Unit Error")
	th.CheckF(t, a.Value, 1.6, "Value Error")
}

func TestRadian(t *testing.T) {
	cases := make(map[Angle]Angle)
	cases[NewAngle(Radian, 1.1)] = NewAngle(Radian, 1.1)
	cases[NewAngle(Degree, 180.)] = NewAngle(Radian, math.Pi)
	cases[NewAngle(Hour, 12.0)] = NewAngle(Radian, math.Pi)
	cases[NewAngle(MilliRadian, 1.1)] = NewAngle(Radian, 0.0011)
	cases[NewAngle(ArcMinute, 1.1)] = NewAngle(Radian, 0.00031997702953229375)
	cases[NewAngle(ArcSecond, 66.0)] = NewAngle(Radian, 0.00031997702953229375)

	for in, ex := range cases {
		g := in.Radian()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Value, ex.Value, "Value Error")
	}
}

func TestMilliRadian(t *testing.T) {
	cases := make(map[Angle]Angle)
	cases[NewAngle(Radian, 1.1)] = NewAngle(MilliRadian, 1100)
	cases[NewAngle(Degree, 180.)] = NewAngle(MilliRadian, 3141.592653589793)
	cases[NewAngle(Hour, 12.0)] = NewAngle(MilliRadian, 3141.592653589793)
	cases[NewAngle(MilliRadian, 1.1)] = NewAngle(MilliRadian, 1.1)
	cases[NewAngle(ArcMinute, 1.1)] = NewAngle(MilliRadian, 0.31997702953229375)
	cases[NewAngle(ArcSecond, 66.0)] = NewAngle(MilliRadian, 0.31997702953229375)

	for in, ex := range cases {
		g := in.MilliRadian()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Value, ex.Value, "Value Error")
	}
}

func TestDegree(t *testing.T) {
	cases := make(map[Angle]Angle)
	cases[NewAngle(Radian, math.Pi)] = NewAngle(Degree, 180.)
	cases[NewAngle(Degree, 180.)] = NewAngle(Degree, 180.)
	cases[NewAngle(Hour, 12.0)] = NewAngle(Degree, 180.)
	cases[NewAngle(MilliRadian, 1.1)] = NewAngle(Degree, 0.06302535746439056)
	cases[NewAngle(ArcMinute, 1.1)] = NewAngle(Degree, 0.018333333333333333)
	cases[NewAngle(ArcMinute, -1.1)] = NewAngle(Degree, -0.018333333333333333)
	cases[NewAngle(ArcSecond, 66.0)] = NewAngle(Degree, 0.018333333333333333)
	cases[NewAngle(ArcSecond, -66.0)] = NewAngle(Degree, -0.018333333333333333)

	for in, ex := range cases {
		g := in.Degree()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Value, ex.Value, "Value Error")
	}
}

func TestArcMinute(t *testing.T) {
	cases := make(map[Angle]Angle)
	cases[NewAngle(Radian, math.Pi)] = NewAngle(ArcMinute, 10800.0)
	cases[NewAngle(Degree, 180.)] = NewAngle(ArcMinute, 10800.0)
	cases[NewAngle(Hour, 12.0)] = NewAngle(ArcMinute, 10800.0)
	cases[NewAngle(MilliRadian, 1.1)] = NewAngle(ArcMinute, 3.7815214478634336)
	cases[NewAngle(ArcMinute, 1.1)] = NewAngle(ArcMinute, 1.1)
	cases[NewAngle(ArcMinute, -1.1)] = NewAngle(ArcMinute, -1.1)
	cases[NewAngle(ArcSecond, 66.0)] = NewAngle(ArcMinute, 1.1)
	cases[NewAngle(ArcSecond, -66.0)] = NewAngle(ArcMinute, -1.1)

	for in, ex := range cases {
		g := in.ArcMinute()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Value, ex.Value, "Value Error")
	}
}

func TestHour(t *testing.T) {
	cases := make(map[Angle]Angle)
	cases[NewAngle(Radian, math.Pi)] = NewAngle(Hour, 12.0)
	cases[NewAngle(Degree, 180.)] = NewAngle(Hour, 12.0)
	cases[NewAngle(Hour, 12.0)] = NewAngle(Hour, 12.0)
	cases[NewAngle(MilliRadian, 1.1)] = NewAngle(Hour, 0.004201690497626037)
	cases[NewAngle(ArcMinute, 1.1)] = NewAngle(Hour, 0.0012222222222222222)
	cases[NewAngle(ArcMinute, -1.1)] = NewAngle(Hour, -0.0012222222222222222)
	cases[NewAngle(ArcSecond, 66.0)] = NewAngle(Hour, 0.0012222222222222222)
	cases[NewAngle(ArcSecond, -66.0)] = NewAngle(Hour, -0.0012222222222222222)

	for in, ex := range cases {
		g := in.Hour()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Value, ex.Value, "Value Error")
	}
}

// Test ArcSecond
func TestArcSecond(t *testing.T) {
	cases := make(map[Angle]Angle)
	cases[NewAngle(Radian, math.Pi)] = NewAngle(ArcSecond, 648000.0)
	cases[NewAngle(Degree, 180.)] = NewAngle(ArcSecond, 648000.0)
	cases[NewAngle(Hour, 12.0)] = NewAngle(ArcSecond, 648000.0)
	cases[NewAngle(MilliRadian, 1.1)] = NewAngle(ArcSecond, 226.89128687180602)
	cases[NewAngle(ArcMinute, 1.1)] = NewAngle(ArcSecond, 66.0)
	cases[NewAngle(ArcMinute, -1.1)] = NewAngle(ArcSecond, -66.0)
	cases[NewAngle(ArcSecond, 66.0)] = NewAngle(ArcSecond, 66.0)
	cases[NewAngle(ArcSecond, -66.0)] = NewAngle(ArcSecond, -66.0)

	for in, ex := range cases {
		g := in.ArcSecond()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Value, ex.Value, "Value Error")
	}
}

// Test MilliArcSecond
func TestMilliArcSecond(t *testing.T) {
	cases := make(map[Angle]Angle)
	cases[NewAngle(Radian, math.Pi)] = NewAngle(MilliArcSecond, 648000000.)
	cases[NewAngle(Degree, 180.)] = NewAngle(MilliArcSecond, 648000000.)
	cases[NewAngle(Hour, 12.0)] = NewAngle(MilliArcSecond, 648000000.)
	cases[NewAngle(MilliRadian, 1.1)] = NewAngle(MilliArcSecond, 226891.28687180602)
	cases[NewAngle(ArcMinute, 1.1)] = NewAngle(MilliArcSecond, 66000.)
	cases[NewAngle(ArcMinute, -1.1)] = NewAngle(MilliArcSecond, -66000.)
	cases[NewAngle(ArcSecond, 66.0)] = NewAngle(MilliArcSecond, 66000.)
	cases[NewAngle(ArcSecond, -66.0)] = NewAngle(MilliArcSecond, -66000.)

	for in, ex := range cases {
		g := in.MilliArcSecond()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Value, ex.Value, "Value Error")
	}
}

// Test SexagesimalDMS
func TestSexagesimalDMS(t *testing.T) {
	cases := make(map[Angle]string)
	cases[NewAngle(Radian, math.Pi)] = "180:00:00.0000"
	cases[NewAngle(Degree, 180.)] = "180:00:00.0000"
	cases[NewAngle(Hour, 12.0)] = "180:00:00.0000"
	cases[NewAngle(MilliRadian, 1.1)] = "00:03:46.8913"
	cases[NewAngle(MilliRadian, -1.1)] = "-00:03:46.8913"
	cases[NewAngle(ArcMinute, 1.1)] = "00:01:06.0000"
	cases[NewAngle(ArcMinute, -1.1)] = "-00:01:06.0000"
	cases[NewAngle(ArcSecond, 66.0)] = "00:01:06.0000"
	cases[NewAngle(ArcSecond, -66.0)] = "-00:01:06.0000"

	for in, ex := range cases {
		g := in.SexagesimalDMS()
		th.CheckS(t, g, ex, "Value Error")
	}
}

// Test SexagesimalHMS
func TestSexagesimalHMS(t *testing.T) {
	cases := make(map[Angle]string)
	cases[NewAngle(Radian, math.Pi)] = "12:00:00.0000"
	cases[NewAngle(Degree, 180.)] = "12:00:00.0000"
	cases[NewAngle(Hour, 12.0)] = "12:00:00.0000"
	cases[NewAngle(MilliRadian, 1.1)] = "00:00:15.1261"

	cases[NewAngle(MilliRadian, -1.1)] = "-00:00:15.1261"

	cases[NewAngle(ArcMinute, 1.1)] = "00:00:04.4000"
	cases[NewAngle(ArcMinute, -1.1)] = "-00:00:04.4000"
	cases[NewAngle(ArcSecond, 66.0)] = "00:00:04.4000"
	cases[NewAngle(ArcSecond, -66.0)] = "-00:00:04.4000"

	for in, ex := range cases {
		g := in.SexagesimalHMS()
		th.CheckS(t, g, ex, "Value Error")
	}
}

// Test Radian receiver for DMS
func TestRadianReceiver(t *testing.T) {
	cases := make(map[DMS]Angle)
	cases[DMS{1.0, 1.0, 1.1}] = NewAngle(Radian, 0.017749513679101225)
	cases[DMS{-1.0, -1.0, -1.1}] = NewAngle(Radian, -0.017749513679101225)
	cases[DMS{-1.0, 1.0, 1.1}] = NewAngle(Radian, -0.017749513679101225)
	cases[DMS{1.0, 90.0, 1.1}] = NewAngle(Radian, 0.04363856425035044)
	for in, ex := range cases {
		g := in.Angle()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Radian().Value, ex.Value, "Value Error")
	}
}

// Test MilliRadian receiver for DMS
func TestMilliRadianReceiver(t *testing.T) {
	cases := make(map[DMS]Angle)
	cases[DMS{1.0, 1.0, 1.1}] = NewAngle(MilliRadian, 17.749513679101224)
	cases[DMS{-1.0, -1.0, -1.1}] = NewAngle(MilliRadian, -17.749513679101224)
	cases[DMS{-1.0, 1.0, 1.1}] = NewAngle(MilliRadian, -17.749513679101224)
	cases[DMS{1.0, 90.0, 1.1}] = NewAngle(MilliRadian, 43.63856425035044)
	for in, ex := range cases {
		g := in.Angle().MilliRadian()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.MilliRadian().Value, ex.Value, "Value Error")
	}
}

// Test Degree receiver for DMS
func TestDegreeReceiver(t *testing.T) {
	cases := make(map[DMS]Angle)
	cases[DMS{1.0, 1.0, 1.1}] = NewAngle(Degree, 1.0169722222222224)
	cases[DMS{-1.0, -1.0, -1.1}] = NewAngle(Degree, -1.0169722222222224)
	cases[DMS{-1.0, 1.0, 1.1}] = NewAngle(Degree, -1.0169722222222224)
	cases[DMS{1.0, 90.0, 1.1}] = NewAngle(Degree, 2.5003055555555553)
	for in, ex := range cases {
		g := in.Angle().Degree()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Degree().Value, ex.Value, "Value Error")
	}
}

// Test ArcMinute receiver for DMS
func TestArcMinuteReceiver(t *testing.T) {
	cases := make(map[DMS]Angle)
	cases[DMS{1.0, 1.0, 1.1}] = NewAngle(ArcMinute, 61.018333333333345)
	cases[DMS{-1.0, -1.0, -1.1}] = NewAngle(ArcMinute, -61.018333333333345)
	cases[DMS{-1.0, 1.0, 1.1}] = NewAngle(ArcMinute, -61.018333333333345)
	cases[DMS{1.0, 90.0, 1.1}] = NewAngle(ArcMinute, 150.01833333333332)
	for in, ex := range cases {
		g := in.Angle().ArcMinute()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.ArcMinute().Value, ex.Value, "Value Error")
	}
}

// Test ArcSecond receiver for DMS
func TestArcSecondReceiver(t *testing.T) {
	cases := make(map[DMS]Angle)
	cases[DMS{1.0, 1.0, 1.1}] = NewAngle(ArcSecond, 3661.100000000001)
	cases[DMS{-1.0, -1.0, -1.1}] = NewAngle(ArcSecond, -3661.100000000001)
	cases[DMS{-1.0, 1.0, 1.1}] = NewAngle(ArcSecond, -3661.100000000001)
	cases[DMS{1.0, 90.0, 1.1}] = NewAngle(ArcSecond, 9001.099999999999)
	for in, ex := range cases {
		g := in.Angle().ArcSecond()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.ArcSecond().Value, ex.Value, "Value Error")
	}
}

// Test MilliArcSecond receiver for DMS
func TestMilliArcSecondReceiver(t *testing.T) {
	cases := make(map[DMS]Angle)
	cases[DMS{1.0, 1.0, 1.1}] = NewAngle(MilliArcSecond, 3661100.000000001)
	cases[DMS{-1.0, -1.0, -1.1}] = NewAngle(MilliArcSecond, -3661100.000000001)
	cases[DMS{-1.0, 1.0, 1.1}] = NewAngle(MilliArcSecond, -3661100.000000001)
	cases[DMS{1.0, 90.0, 1.1}] = NewAngle(MilliArcSecond, 9001099.999999999)
	for in, ex := range cases {
		g := in.Angle().MilliArcSecond()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckFT(t, g.MilliArcSecond().Value, ex.Value, tol, "Value Error")
	}
}

// Test HMS receiver for DMS
func TestHMSReceiver(t *testing.T) {
	cases := make(map[DMS]HMS)
	cases[DMS{15.0, 1.0, 1.1}] = HMS{1.0, 0.0, 4.07333333333364}
	cases[DMS{-15.0, -1.0, -1.1}] = HMS{-1.0, 0.0, -4.07333333333364}
	cases[DMS{-15.0, 1.0, 1.1}] = HMS{-1.0, 0.0, -4.07333333333364}
	cases[DMS{14.0, 90.0, 1.1}] = HMS{1.0, 2.0, 0.07333333333338032}
	cases[DMS{-14.0, 90.0, 1.1}] = HMS{-1.0, -2.0, -0.07333333333338032}
	for in, ex := range cases {
		g := in.HMS()
		th.CheckF(t, g.Hr, ex.Hr, "Hr Error")
		th.CheckF(t, g.Min, ex.Min, "Min Error")
		th.CheckF(t, g.Sec, ex.Sec, "Sec Error")
	}
}

// Test Radian receiver for HMS
func TestRadianHMSReceiver(t *testing.T) {
	cases := make(map[HMS]Angle)
	cases[HMS{12.0, 1.0, 1.1}] = NewAngle(Radian, 3.1460359709771626)
	cases[HMS{-12.0, -1.0, -1.1}] = NewAngle(Radian, -3.1460359709771626)
	cases[HMS{-12.0, 1.0, 1.1}] = NewAngle(Radian, -3.1460359709771626)
	cases[HMS{11.0, 90.0, 1.1}] = NewAngle(Radian, 3.2725723417467516)
	for in, ex := range cases {
		g := in.Angle()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Radian().Value, ex.Value, "Value Error")
	}
}

// Test MilliRadian receiver for HMS
func TestMilliRadianHMSReceiver(t *testing.T) {
	cases := make(map[HMS]Angle)
	cases[HMS{12.0, 1.0, 1.1}] = NewAngle(MilliRadian, 3146.0359709771626)
	cases[HMS{-12.0, -1.0, -1.1}] = NewAngle(MilliRadian, -3146.0359709771626)
	cases[HMS{-12.0, 1.0, 1.1}] = NewAngle(MilliRadian, -3146.0359709771626)
	cases[HMS{11.0, 90.0, 1.1}] = NewAngle(MilliRadian, 3272.5723417467516)
	for in, ex := range cases {
		g := in.Angle().MilliRadian()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.MilliRadian().Value, ex.Value, "Value Error")
	}
}

// Test Degree receiver for HMS
func TestDegreeHMSReceiver(t *testing.T) {
	cases := make(map[HMS]Angle)
	cases[HMS{12.0, 1.0, 1.1}] = NewAngle(Degree, 180.25458333333336)
	cases[HMS{-12.0, -1.0, -1.1}] = NewAngle(Degree, -180.25458333333336)
	cases[HMS{-12.0, 1.0, 1.1}] = NewAngle(Degree, -180.25458333333336)
	cases[HMS{11.0, 90.0, 1.1}] = NewAngle(Degree, 187.50458333333336)
	for in, ex := range cases {
		g := in.Angle().Degree()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.Degree().Value, ex.Value, "Value Error")
	}
}

// Test ArcMinute receiver for HMS
func TestArcMinuteHMSReceiver(t *testing.T) {
	cases := make(map[HMS]Angle)
	cases[HMS{12.0, 1.0, 1.1}] = NewAngle(ArcMinute, 10815.275000000001)
	cases[HMS{-12.0, -1.0, -1.1}] = NewAngle(ArcMinute, -10815.275000000001)
	cases[HMS{-12.0, 1.0, 1.1}] = NewAngle(ArcMinute, -10815.275000000001)
	cases[HMS{11.0, 90.0, 1.1}] = NewAngle(ArcMinute, 11250.275000000001)
	for in, ex := range cases {
		g := in.Angle().ArcMinute()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.ArcMinute().Value, ex.Value, "Value Error")
	}
}

// Test ArcSecond receiver for HMS
func TestArcSecondHMSReceiver(t *testing.T) {
	cases := make(map[HMS]Angle)
	cases[HMS{12.0, 1.0, 1.1}] = NewAngle(ArcSecond, 648916.5000000001)
	cases[HMS{-12.0, -1.0, -1.1}] = NewAngle(ArcSecond, -648916.5000000001)
	cases[HMS{-12.0, 1.0, 1.1}] = NewAngle(ArcSecond, -648916.5000000001)
	cases[HMS{11.0, 90.0, 1.1}] = NewAngle(ArcSecond, 675016.5000000001)
	for in, ex := range cases {
		g := in.Angle().ArcSecond()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckF(t, g.ArcSecond().Value, ex.Value, "Value Error")
	}
}

// Test MilliArcSecond receiver for HMS
func TestMilliArcSecondHMSReceiver(t *testing.T) {
	cases := make(map[HMS]Angle)
	cases[HMS{12.0, 1.0, 1.1}] = NewAngle(MilliArcSecond, 648916500.0000001)
	cases[HMS{-12.0, -1.0, -1.1}] = NewAngle(MilliArcSecond, -648916500.0000001)
	cases[HMS{-12.0, 1.0, 1.1}] = NewAngle(MilliArcSecond, -648916500.0000001)
	cases[HMS{11.0, 90.0, 1.1}] = NewAngle(MilliArcSecond, 675016500.0000001)
	for in, ex := range cases {
		g := in.Angle().MilliArcSecond()
		th.CheckS(t, g.UnitString(), ex.UnitString(), "Unit Error")
		th.CheckFT(t, g.MilliArcSecond().Value, ex.Value, tol, "Value Error")
	}
}

// Test DMS receiver for HMS
func TestDMSReceiver(t *testing.T) {
	cases := make(map[HMS]DMS)
	cases[HMS{12.0, 1.0, 1.1}] = DMS{180.0, 15.0, 16.50000000008731}
	cases[HMS{-12.0, -1.0, -1.1}] = DMS{-180.0, -15.0, -16.50000000008731}
	cases[HMS{-12.0, 1.0, 1.1}] = DMS{-180.0, -15.0, -16.50000000008731}
	cases[HMS{11.0, 90.0, 1.1}] = DMS{187.0, 30.0, 16.50000000008731}
	cases[HMS{-11.0, 90.0, 1.1}] = DMS{-187.0, -30.0, -16.50000000008731}
	for in, ex := range cases {
		g := in.DMS()
		th.CheckF(t, g.Deg, ex.Deg, "Deg Error")
		th.CheckF(t, g.Min, ex.Min, "Min Error")
		th.CheckF(t, g.Sec, ex.Sec, "Sec Error")
	}
}

func TestAngleModuloN(t *testing.T) {
	m := ModuloN(360. * 4)
	a := m(360. * 2.)
	th.CheckFT(t, a, 360.*2., 1e-9, "Value Error")
}

func TestConstrain(t *testing.T) {
	in := []Angle{
		NewAngle(Degree, 45.0),
		NewAngle(Degree, 360.0),
		NewAngle(Degree, -30.0),
		NewAngle(Degree, -100.0)}
	minA := NewAngle(Degree, -90.0)
	maxA := NewAngle(Degree, 359.0)
	ex := []Angle{
		NewAngle(Degree, 45.0),
		maxA,
		NewAngle(Degree, -30.0),
		minA}
	for idx, a := range in {
		got := a
		got.Constrain(minA, maxA)
		th.CheckFT(t, got.Degree().Value, ex[idx].Degree().Value, 1e-9, "Value Error")
	}

}

func TestMap(t *testing.T) {
	in := []Angle{
		NewAngle(Degree, 45.0),
		NewAngle(Degree, 45.0),
		NewAngle(Degree, 359.9),
		NewAngle(Degree, 359.9),
		NewAngle(Degree, 0.0),
		NewAngle(Degree, 0.0),
		NewAngle(Degree, 350.0),
		NewAngle(Degree, 350.0),
		NewAngle(Degree, 350.0)}
	wc := []int32{0, 1, 0, 1, 0, 1, -1, 0, 1}
	ex := []Angle{
		NewAngle(Degree, 45.0),
		NewAngle(Degree, 405.0),
		NewAngle(Degree, 359.9),
		NewAngle(Degree, 719.9),
		NewAngle(Degree, 0.0),
		NewAngle(Degree, 360.0),
		NewAngle(Degree, -10.0),
		NewAngle(Degree, 350.0),
		NewAngle(Degree, 710.0)}

	for idx, a := range in {
		got := Map(a, wc[idx])
		msg := fmt.Sprintf("Value Error for input: %6.3f, wc= %d", a.Degree().Value, wc[idx])
		th.CheckFT(t, got.Degree().Value, ex[idx].Degree().Value, 1e-9, msg)
	}
}

func TestClosest(t *testing.T) {
	maxWrapCount := int32(5)
	// current angle
	ca := []Angle{
		NewAngle(Degree, 350.0),
		NewAngle(Degree, 350.0),
		NewAngle(Degree, 370.0),
		NewAngle(Degree, 370.0),
		NewAngle(Degree, 10.0),
		NewAngle(Degree, -10.0),
		NewAngle(Degree, 1790.0)}
	// desired angle
	da := []Angle{
		NewAngle(Degree, 45.0),
		NewAngle(Degree, 300.0),
		NewAngle(Degree, 20.0),
		NewAngle(Degree, 350.0),
		NewAngle(Degree, 350.0),
		NewAngle(Degree, 10.0),
		NewAngle(Degree, 45.0)}
	// wrap count
	wc := []int32{0, 0, 1, 1, 0, -1, 4}
	// expected angle
	ex := []Angle{
		NewAngle(Degree, 405.0),
		NewAngle(Degree, 300.0),
		NewAngle(Degree, 380.0),
		NewAngle(Degree, 350.0),
		NewAngle(Degree, -10.0),
		NewAngle(Degree, 10.0),
		NewAngle(Degree, 1485.0)}
	for idx, a := range da {
		got := Closest(ca[idx], a, wc[idx], maxWrapCount)
		th.CheckFT(t, got.Degree().Value, ex[idx].Degree().Value, 1e-1, "Value Error")
	}
}

func TestSin(t *testing.T) {
	iA := []Angle{
		NewAngle(Degree, 0.0),
		NewAngle(Degree, 90.0),
		NewAngle(Degree, 30.0),
		NewAngle(Degree, -30.0)}

	ev := []float64{
		0.0,
		1.0,
		0.5,
		-0.5}
	tol := 1e-5
	for idx, a := range iA {
		got := a.Sin()
		th.CheckFT(t, got, ev[idx], tol, "Error Value")
	}
}

func TestNewHMS(t *testing.T) {
	in := []string{"01:02:03.123", "+01:02:03.123", "-01:02:03.123"}
	exHr := 1.0
	exMin := 2.0
	exSec := 3.123
	for idx, i := range in {
		hms, err := NewHMS(i)
		if err != nil {
			fmt.Sprintf("Error from NewHMS: %v", err)
			t.Fail()
		}

		if idx < 2 {
			th.CheckFT(t, hms.Hr, exHr, 1.e-6, "Value Error")
			th.CheckFT(t, hms.Min, exMin, 1.e-6, "Value Error")
			th.CheckFT(t, hms.Sec, exSec, 1.e-6, "Value Error")
		} else {
			th.CheckFT(t, hms.Hr, -exHr, 1.e-6, "Value Error")
			th.CheckFT(t, hms.Min, -exMin, 1.e-6, "Value Error")
			th.CheckFT(t, hms.Sec, -exSec, 1.e-6, "Value Error")
		}
	}
}

func TestNewDMS(t *testing.T) {
	in := []string{"01:02:03.123", "+01:02:03.123", "-01:02:03.123"}
	exDeg := 1.0
	exMin := 2.0
	exSec := 3.123
	for idx, i := range in {
		dms, err := NewDMS(i)
		if err != nil {
			fmt.Sprintf("Error from NewDMS: %v", err)
			t.Fail()
		}

		if idx < 2 {
			th.CheckFT(t, dms.Deg, exDeg, 1.e-6, "Value Error")
			th.CheckFT(t, dms.Min, exMin, 1.e-6, "Value Error")
			th.CheckFT(t, dms.Sec, exSec, 1.e-6, "Value Error")
		} else {
			th.CheckFT(t, dms.Deg, -exDeg, 1.e-6, "Value Error")
			th.CheckFT(t, dms.Min, -exMin, 1.e-6, "Value Error")
			th.CheckFT(t, dms.Sec, -exSec, 1.e-6, "Value Error")
		}
	}
}

func TestPow(t *testing.T) {
	a := NewAngle(Radian, 2.0)
	b := a.Pow(2.0)
	exp := 4.0
	th.CheckFT(t, exp, b.Radian().Value, tol, "Value Error")

	b = a.Pow(0.5)
	exp = 1.4142135623730951
	th.CheckFT(t, exp, b.Radian().Value, tol, "Value Error")

}
