package features

import (
	"testing"

	"github.com/nishantb06/ray_tracing_challenge_golang/features"
)

func TestMatrix(t *testing.T) {
	m := features.Matrix_(4, 4)
	m.FillValues(1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5)
	if m.Get(0, 0) != 1 {
		t.Error("m[0][0] should be 1")
	}
	if m.Get(0, 3) != 4 {
		t.Error("m[0][3] should be 4")
	}
	if m.Get(1, 0) != 5.5 {
		t.Error("m[1][0] should be 5.5")
	}
	if m.Get(1, 2) != 7.5 {
		t.Error("m[1][2] should be 7.5")
	}
	if m.Get(2, 2) != 11 {
		t.Error("m[2][2] should be 11")
	}
	if m.Get(3, 0) != 13.5 {
		t.Error("m[3][0] should be 13.5")
	}
	if m.Get(3, 2) != 15.5 {
		t.Error("m[3][2] should be 15.5")
	}

	m2 := features.Matrix_(2, 2)
	m2.FillValues(-3, 5, 1, -2)
	if m2.Get(0, 0) != -3 {
		t.Error("m2[0][0] should be -3")
	}
	if m2.Get(0, 1) != 5 {
		t.Error("m2[0][1] should be 5")
	}
	if m2.Get(1, 0) != 1 {
		t.Error("m2[1][0] should be 1")
	}
	if m2.Get(1, 1) != -2 {
		t.Error("m2[1][1] should be -2")
	}

	// a 3x3 matrix ought to be representable
	m3 := features.Matrix_(3, 3)
	m3.FillValues(-3, 5, 0, 1, -2, -7, 0, 1, 1)
	if m3.Get(0, 0) != -3 {
		t.Error("m3[0][0] should be -3")
	}
	if m3.Get(1, 1) != -2 {
		t.Error("m3[1][1] should be -2")
	}
	if m3.Get(2, 1) != 1 {
		t.Error("m3[2][1] should be 1")
	}

	// Matrix equality with identical matrices
	m4 := features.Matrix_(4, 4)
	m4.FillValues(1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5)
	if !m.Equals(m4) {
		t.Error("m should be equal to m4")
	}
	m5 := features.Matrix_(4, 4)
	m5.FillValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	if m.Equals(m5) {
		t.Error("m should not be equal to m5")
	}

}
