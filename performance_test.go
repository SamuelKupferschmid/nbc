package nbc

import (
	"math"
	"testing"
)

func TestPerformanceF1(t *testing.T) {

	if (&Performance{
		1,
		1,
	}).F1() != 1 {
		t.Fail()
	}

	if !math.IsNaN((&Performance{
		0,
		0,
	}).F1()) {
		t.Fail()
	}
}
