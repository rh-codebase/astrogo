package astrounit

import (
	"fmt"
	"math"
	"time"
)

// deGlitch is a closure for deglitch erroneous encoder readings
func DeGlitch(a Angle, ti time.Time, maxr AngularRate) func(Angle, time.Time) Angle {
	maxrateV := maxr.Degree().Value
	prevAngle := a
	prevTime := ti
	return func(cura Angle, t time.Time) Angle {
		diffA := cura.Sub(prevAngle)
		diffT := t.Sub(prevTime)
		rate := NewAngularRate(diffA, diffT)
		ratev := rate.Degree().Value
		//fmt.Println("pT: ", prevTime, " pA: ", prevAngle.Degree().Value, " ct: ", t, " ca: ", cura.Degree().Value, " r: ", ratev)

		if math.Abs(ratev) > maxrateV {
			fmt.Println("DeGlitching: time:", t, " cura: ", cura.Degree().Value, " prevA: ", prevAngle.Degree().Value, " rate: ", ratev)
			return NewAngle(Degree, prevAngle.Degree().Value+maxrateV*diffT.Seconds())
		} else {
			prevAngle = cura
			prevTime = t
			return cura
		}
	}
}
