package astrotime

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIsoNow(t *testing.T) {
	fmt.Println(IsoNow())
}

func TestIsoNowMilli(t *testing.T) {
	fmt.Println(IsoNowMilli())
}

func TestIsoFormat(t *testing.T) {
	ti := time.Now()
	fmt.Println(IsoFormat(ti, RFC3339deci))
	fmt.Println(IsoFormat(ti, RFC3339centi))
	fmt.Println(IsoFormat(ti, RFC3339milli))
	fmt.Println(IsoFormat(ti, RFC3339micro))
	fmt.Println(IsoFormat(ti, RFC3339nano))
}

func TestParseTime(t *testing.T) {
	myt := "2024-11-04T14:22:00.000-07:00"
	ti, err := ParseTime(myt)
	if err != nil {
		fmt.Println("Error ParseTime: ", err)
		t.Fail()
	}
	fmt.Println(ti)
}

func TestParseLocalTimeToUTC(t *testing.T) {
	myt := "2024-11-04T14:22:00.000-07:00"
	ti, err := ParseLocalTimeToUTC(myt)
	if err != nil {
		fmt.Println("Error ParseLocalTimeToUTC: ", err)
		t.Fail()
	}
	fmt.Println(ti)
}

func TestSchedule(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	f := func(t time.Time) { fmt.Println(t) }
	go Schedule(ctx, time.Second, 0, f)
	time.Sleep(5 * time.Second)
	cancel()
}
