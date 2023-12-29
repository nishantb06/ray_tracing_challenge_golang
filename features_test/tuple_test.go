package features

import (
	"fmt"
	"testing"

	"github.com/nishantb06/ray_tracing_challenge_golang/features"
)

func TestGetTuple(t *testing.T) {
	fmt.Println("Hello, World!")
	p, error := features.GetTuple("point", 1, 2, 3)
	if error != nil {
		t.Error(error)
	}
	if p.GetProperty() != "point" {
		t.Error("Point property should be point")
	}
	if p.GetW() != 1 {
		t.Error("coordinate w should be 1")
	}

	v, error := features.GetTuple("vector", 1, 2, 3)
	if error != nil {
		t.Error(error)
	}
	if v.GetProperty() != "vector" {
		t.Error("Vector property should be vector")
	}
	if v.GetW() != 0 {
		t.Error("coordinate w should be 0")
	}

	p1 := features.Point(1, 2, 3)
	if p1.GetProperty() != "point" {
		t.Error("Point property should be point")
	}
	if p1.GetW() != 1 {
		t.Error("coordinate w should be 1")
	}

	v1 := features.Vector(1, 2, 3)
	if v1.GetProperty() != "vector" {
		t.Error("Vector property should be vector")
	}
	if v1.GetW() != 0 {
		t.Error("coordinate w should be 0")
	}
}

func TestAdd(t *testing.T) {
	a1 := features.Point(3, -2, 5)
	a2 := features.Vector(-2, 3, 1)
	a3 := features.Add(a1, a2)
	if a3.GetProperty() != "point" {
		t.Error("Point property should be point")
	}
	if a3.GetW() != 1 {
		t.Error("coordinate w should be 1")
	}
	if a3.GetX() != 1 {
		t.Error("coordinate x should be 1")
	}
	if a3.GetY() != 1 {
		t.Error("coordinate y should be 1")
	}
	if a3.GetZ() != 6 {
		t.Error("coordinate z should be 6")
	}
}
