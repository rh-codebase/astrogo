package astrounit

import (
	"fmt"
	"testing"
	"time"

	th "github.com/rh-codebase/genutilsgo"
)

const (
	tol = 1e-6
)

func TestModulo24(t *testing.T) {
	mh := Modulo24(25.0)
	th.CheckF(t, mh, 1, "Error:")

	mh = Modulo24(-1.0)
	th.CheckF(t, mh, 23, "Error:")

	mh = Modulo24(25.1)
	th.CheckFT(t, mh, 1.1, .0000000000001, "Error:")

	mh = Modulo24(49.1)
	th.CheckFT(t, mh, 1.1, .0000000000001, "Error:")

	mh = Modulo24(-25.1)
	th.CheckFT(t, mh, 22.9, .0000000000001, "Error:")

	mh = Modulo24(-49.1)
	th.CheckFT(t, mh, 22.9, .0000000000001, "Error:")

	mh = Modulo24(23.9)
	th.CheckFT(t, mh, 23.9, .0000000000001, "Error:")

	mh = Modulo24(-23.9)
	th.CheckFT(t, mh, 0.1, .0000000000001, "Error:")

	mh = Modulo24(0.0)
	th.CheckFT(t, mh, 0.0, .0000000000001, "Error:")

	mh = Modulo24(-0.0)
	th.CheckFT(t, mh, 0.0, .0000000000001, "Error:")

	mh = Modulo24(24.0)
	th.CheckFT(t, mh, 0.0, .0000000000001, "Error:")

	mh = Modulo24(-24.0)
	th.CheckFT(t, mh, 0.0, .0000000000001, "Error:")
}

func TestModuloN(t *testing.T) {
	m := ModuloN(24.0)
	mh := m(25.0)
	th.CheckF(t, mh, 1, "Error:")

	mh = m(-1.0)
	th.CheckF(t, mh, 23, "Error:")

	mh = m(25.1)
	th.CheckFT(t, mh, 1.1, .0000000000001, "Error:")

	m = ModuloN(360. * 4)
	mh = m(25.0)
	th.CheckF(t, mh, 25, "Error:")

	mh = m(-1.0)
	th.CheckF(t, mh, 360.*4.-1., "Error:")

	mh = m(366.)
	th.CheckFT(t, mh, 366., .0000000000001, "Error:")

	mh = m(360.*4. + 5.)
	th.CheckFT(t, mh, 5., .0000000000001, "Error:")
}

func TestSecondsAfterBoundry(t *testing.T) {
	boundry := []time.Duration{100 * time.Millisecond,
		10 * time.Millisecond,
		time.Millisecond,
		time.Second}
	exdt := []float64{0.0720561416, 0.0020561416, .0000561416, .1720561416}
	//ct := time.Now()
	ct := time.Unix(1720561416, 172056141)

	for idx, b := range boundry {
		dt := SecondsAfterBoundry(ct, b)
		th.CheckFT(t, dt, exdt[idx], tol, "Error Value:")
	}
}

func TestWrapCounter(t *testing.T) {
	encA := NewAngle(Degree, 359.9)
	iwc := int32(1)
	wrapc := WrapCounter(encA, iwc)
	time.Sleep(100 * time.Millisecond)
	wc := wrapc(NewAngle(Degree, 0.))
	if wc != 2 {
		fmt.Println("1: Value Error: Got: ", wc, " Expected 2")
		t.Fail()
	}
	time.Sleep(100 * time.Millisecond)
	wc = wrapc(NewAngle(Degree, 359.9))
	if wc != 1 {
		fmt.Println("2: Value Error: Got: ", wc, " Expected 1")
		t.Fail()
	}
	time.Sleep(100 * time.Millisecond)
	wc = wrapc(NewAngle(Degree, 359.8))
	if wc != 1 {
		fmt.Println("3: Value Error: Got: ", wc, " Expected 1")
		t.Fail()
	}
	time.Sleep(100 * time.Millisecond)
	wc = wrapc(NewAngle(Degree, 0.2))
	if wc != 1 {
		fmt.Println("4: Value Error: Got: ", wc, " Expected 1")
		t.Fail()
	}
	time.Sleep(100 * time.Millisecond)
	wc = wrapc(NewAngle(Degree, 359.9))
	if wc != 1 {
		fmt.Println("5: Value Error: Got: ", wc, " Expected 1")
		t.Fail()
	}
	time.Sleep(100 * time.Millisecond)
	wc = wrapc(NewAngle(Degree, 0.))
	if wc != 2 {
		fmt.Println("6: Value Error: Got: ", wc, " Expected 1")
		t.Fail()
	}
}

func TestMapAngle(t *testing.T) {
	encA := NewAngle(Degree, 10.0)
	for wc := -3; wc < 4; wc++ {
		ma := MapAngle(int32(wc), encA)
		expA := NewAngle(Degree, 10.+float64(wc)*360.)
		if !ma.Equal(expA) {
			fmt.Println("Value Error. Got: ", ma.Degree().Value, " Expected: ",
				expA.Degree().Value)
			t.Fail()
		}
	}
}
