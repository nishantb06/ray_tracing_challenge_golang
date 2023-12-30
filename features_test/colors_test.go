package features

import (
	"testing"

	"github.com/nishantb06/ray_tracing_challenge_golang/features"
)

// const epsilon = 0.00001

func TestColor(t *testing.T) {
	c := features.Color_(1, 2, 3)
	if c.Red != 1 {
		t.Error("Red should be 1")
	}
	if c.Green != 2 {
		t.Error("Green should be 2")
	}
	if c.Blue != 3 {
		t.Error("Blue should be 3")
	}
}

func TestAddColor(t *testing.T) {
	c1 := features.Color_(1, 2, 3)
	c2 := features.Color_(2, 3, 4)
	c3 := features.AddColor(c1, c2)
	c4 := c1.Add(c2)
	if c3.Red != 3 {
		t.Error("Red should be 3")
	}
	if c3.Green != 5 {
		t.Error("Green should be 5")
	}
	if c3.Blue != 7 {
		t.Error("Blue should be 7")
	}
	if c4.Red != 3 {
		t.Error("Red should be 3")
	}
	if c4.Green != 5 {
		t.Error("Green should be 5")
	}
	if c4.Blue != 7 {
		t.Error("Blue should be 7")
	}
}
