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
	a3, _ := features.Add(a1, a2)
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

func TestSubtract(t *testing.T) {

	a1 := features.Point(3, 2, 1)
	a2 := features.Point(5, 6, 7)
	v1 := features.Vector(5, 6, 7)
	v2 := features.Vector(3, 2, 1)

	// subtracting 2 points,should give a vector
	a3, _ := features.Subtract(a1, a2)
	if a3.GetProperty() != "vector" {
		t.Error("Vector property should be vector")
	}
	if a3.GetW() != 0.0 {
		t.Error("coordinate w should be 0")
	}
	if a3.GetX() != -2 {
		t.Error("coordinate x should be -2")
	}
	if a3.GetY() != -4 {
		t.Error("coordinate y should be -4")
	}
	if a3.GetZ() != -6 {
		t.Error("coordinate z should be -6")
	}

	// subtracting a vector from a point, should give a point
	a4, _ := features.Subtract(a1, v1)
	if a4.GetProperty() != "point" {
		t.Error("Point property should be point")
	}
	if a4.GetW() != 1.0 {
		t.Error("coordinate w should be 1")
	}
	if a4.GetX() != -2 {
		t.Error("coordinate x should be -2")
	}
	if a4.GetY() != -4 {
		t.Error("coordinate y should be -4")
	}
	if a4.GetZ() != -6 {
		t.Error("coordinate z should be -6")
	}

	// subtracting 2 vectors, should give a vector
	a5, _ := features.Subtract(v2, v1)
	if a5.GetProperty() != "vector" {
		t.Error("Vector property should be vector")
	}
	if a5.GetW() != 0.0 {
		t.Error("coordinate w should be 0")
	}
	if a5.GetX() != -2 {
		t.Error("coordinate x should be -2")
	}
	if a5.GetY() != -4 {
		t.Error("coordinate y should be -4")
	}
	if a5.GetZ() != -6 {
		t.Error("coordinate z should be -6")
	}
}

func TestNegate(t *testing.T) {
	a1 := features.Point(1, -2, 3)
	a2, _ := features.Negate(a1)
	if a2.GetProperty() != "point" {
		t.Error("Point property should be point")
	}
	if a2.GetW() != -1.0 {
		t.Error("coordinate w should be 1")
	}
	if a2.GetX() != -1 {
		t.Error("coordinate x should be -1")
	}
	if a2.GetY() != 2 {
		t.Error("coordinate y should be 2")
	}
	if a2.GetZ() != -3 {
		t.Error("coordinate z should be -3")
	}

	v1 := features.Vector(1, -2, 3)
	v2, _ := features.Negate(v1)
	if v2.GetProperty() != "vector" {
		t.Error("Vector property should be vector")
	}
	if v2.GetW() != 0.0 {
		t.Error("coordinate w should be 0")
	}
	if v2.GetX() != -1 {
		t.Error("coordinate x should be -1")
	}
	if v2.GetY() != 2 {
		t.Error("coordinate y should be 2")
	}
	if v2.GetZ() != -3 {
		t.Error("coordinate z should be -3")
	}
}

func TestMultiply(t *testing.T) {
	a1 := features.Point(-2, 3, -4)
	a2, _ := features.Multiply(a1, 3.5)
	if a2.GetProperty() != "point" {
		t.Error("Point property should be point")
	}
	if a2.GetW() != 3.5 {
		t.Error("coordinate w should be 3.5")
	}
	if a2.GetX() != -7 {
		t.Error("coordinate x should be 3.5")
	}
	if a2.GetY() != 10.5 {
		t.Error("coordinate y should be -7")
	}
	if a2.GetZ() != -14.0 {
		t.Error("coordinate z should be 14")
	}

	v1 := features.Vector(1, -2, 3)
	v2, _ := features.Multiply(v1, 3.5)
	if v2.GetProperty() != "vector" {
		t.Error("Vector property should be vector")
	}
	if v2.GetW() != 0.0 {
		t.Error("coordinate w should be 0")
	}
	if v2.GetX() != 3.5 {
		t.Error("coordinate x should be 3.5")
	}
	if v2.GetY() != -7 {
		t.Error("coordinate y should be -7")
	}
	if v2.GetZ() != 10.5 {
		t.Error("coordinate z should be 10.5")
	}

	v3, _ := features.Multiply(v1, 0.5)
	if v3.GetProperty() != "vector" {
		t.Error("Vector property should be vector")
	}
	if v3.GetW() != 0.0 {
		t.Error("coordinate w should be 0")
	}
	if v3.GetX() != 0.5 {
		t.Error("coordinate x should be 0.5")
	}
	if v3.GetY() != -1 {
		t.Error("coordinate y should be -1")
	}
	if v3.GetZ() != 1.5 {
		t.Error("coordinate z should be 1.5")
	}

	a3, _ := features.Multiply(a1, 0.5)
	if a3.GetProperty() != "point" {
		t.Error("Point property should be point")
	}
	if a3.GetW() != 0.5 {
		t.Error("coordinate w should be 0.5")
	}
	if a3.GetX() != -1 {
		t.Error("coordinate x should be -1")
	}
	if a3.GetY() != 1.5 {
		t.Error("coordinate y should be 1.5")
	}
	if a3.GetZ() != -2 {
		t.Error("coordinate z should be -2")
	}
}

func TestDivide(t *testing.T) {
	a1 := features.Point(-2, 3, -4)
	a2, _ := features.Divide(a1, 2)
	if a2.GetProperty() != "point" {
		t.Error("Point property should be point")
	}
	if a2.GetW() != 0.5 {
		t.Error("coordinate w should be 0.5")
	}
	if a2.GetX() != -1 {
		t.Error("coordinate x should be -1")
	}
	if a2.GetY() != 1.5 {
		t.Error("coordinate y should be 1.5")
	}
	if a2.GetZ() != -2 {
		t.Error("coordinate z should be -2")
	}

	v1 := features.Vector(1, -2, 3)
	v2, _ := features.Divide(v1, 2)
	if v2.GetProperty() != "vector" {
		t.Error("Vector property should be vector")
	}
	if v2.GetW() != 0.0 {
		t.Error("coordinate w should be 0")
	}
	if v2.GetX() != 0.5 {
		t.Error("coordinate x should be 0.5")
	}
	if v2.GetY() != -1 {
		t.Error("coordinate y should be -1")
	}
	if v2.GetZ() != 1.5 {
		t.Error("coordinate z should be 1.5")
	}
}
