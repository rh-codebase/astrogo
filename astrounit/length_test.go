package astrounit

import (
	"fmt"
	"testing"
)

func TestNewLength(t *testing.T) {
	l := NewLength(Meter, 123.321)
	if l.Unit != Meter {
		fmt.Println("Fail Unit: Expected Meter, Got ", l.Unit)
		t.Fail()
	}
	if l.Value != 123.321 {
		fmt.Println("Fail Value: Expected 123.321, Got ", l.Value)
		t.Fail()
	}
}
