// Package astrotime contains functions for astronomical time calculations
package astrotime

import (
	"context"
	"time"
)

const (
	JulianDayZero     = float64(2400000.5)
	JulianCentury     = float64(36525.0)
	UnixMJDoffsetDays = float64(40587.) // Day offset between Unix epoch and MJD
	SecondPerMinute   = float64(60.0)
	SecondPerDay      = float64(86400.0)
	SecondPerHour     = float64(3600.0)
	MinutePerDay      = float64(1440.0)
	MinutePerHour     = float64(60.0)
	HourPerDay        = float64(24.0)

	RFC3339deci      = "2006-01-02T15:04:05.9Z07:00"
	RFC3339centi     = "2006-01-02T15:04:05.99Z07:00"
	RFC3339milli     = "2006-01-02T15:04:05.999Z07:00"
	RFC3339micro     = "2006-01-02T15:04:05.999999Z07:00"
	RFC3339nano      = "2006-01-02T15:04:05.999999999Z07:00"
	RFC3339nanoParse = "2006-01-02T15:04:05.999999999-07:00"
)

// MJD2JulianDay return the Julian day for the given MJD
func MJD2JulianDay(mjd float64) float64 {
	return mjd + JulianCentury //JulianDayZero
}

// ToMJD converts a time.Time to MJD.
func ToMJD(t time.Time) float64 {
	return float64(t.Unix())/SecondPerDay + UnixMJDoffsetDays
}

// Turns the current MJD
func MJDNow() float64 {
	t := time.Now()
	return ToMJD(t)
}

// IsoNow returns time as a string in iso8601 format with 0.1 second resolution.
// i.e. YYYY-MM-DDTHH:MM:SS.s-00:07 or YYYY-MM-DDTHH:MM:SS.ssssssZ
func IsoNow() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func IsoNowMilli() string {
	t := time.Now()
	return t.Format(RFC3339milli)
}

// IsoFormat can be used to from to any specification. Helpful formats
// are included in this package: RFC3339deci, RFC3339centi, RFC3339milli,
// RFC3339micro, RFC3339nano
func IsoFormat(t time.Time, format string) string {
	return t.Format(format)
}

// ParseTime parses an ISO3339 time format: 2024-11-04T14:22:00.000-07:00 and
// returns the equivalent time.Time object
func ParseTime(t string) (time.Time, error) {
	return time.Parse(RFC3339nano, t)
}

// ParseLocalTimeUTC parses an ISO3339 time format:
// 2024-11-04T14:22:00.000-07:00 and returns the equivalent UTC
// time.Time object.
func ParseLocalTimeToUTC(t string) (time.Time, error) {
	ti, err := time.Parse(RFC3339nano, t)
	if err != nil {
		return time.Now(), err
	}
	return ti.UTC(), nil
}

// see: https://stackoverflow.com/questions/19549199/golang-implementing-a-cron-executing-tasks-at-a-specific-time
func Schedule(ctx context.Context, p time.Duration, o time.Duration, f func(time.Time)) {
	// Position the first execution
	first := time.Now().Truncate(p).Add(o)
	if first.Before(time.Now()) {
		first = first.Add(p)
	}
	firstC := time.After(first.Sub(time.Now()))

	// Receiving from a nil channel blocks forever
	t := &time.Ticker{C: nil}

	for {
		select {
		case v := <-firstC:
			// The ticker has to be started before f as it can take some time to finish
			t = time.NewTicker(p)
			f(v)
		case v := <-t.C:
			f(v)
		case <-ctx.Done():
			t.Stop()
			return
		}
	}

}
