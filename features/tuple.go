package features

import "fmt"

type Tuple interface {
	GetX() float64
	GetY() float64
	GetZ() float64
	GetW() float64
	setProperty(property string)
	GetProperty() string
}

type PointArray struct {
	x, y, z, w float64
	property   string
}

type VectorArray struct {
	x, y, z, w float64
	property   string
}

func (p *PointArray) setProperty(property string) {
	p.property = property
}

func (p *PointArray) GetProperty() string {
	return p.property
}

func (p *PointArray) GetX() float64 {
	return p.x
}

func (p *PointArray) GetY() float64 {
	return p.y
}

func (p *PointArray) GetZ() float64 {
	return p.z
}

func (p *PointArray) GetW() float64 {
	return p.w
}

func (v *VectorArray) setProperty(property string) {
	v.property = property
}

func (v *VectorArray) GetProperty() string {
	return v.property
}

func (v *VectorArray) GetX() float64 {
	return v.x
}

func (v *VectorArray) GetY() float64 {
	return v.y
}

func (v *VectorArray) GetZ() float64 {
	return v.z
}

func (v *VectorArray) GetW() float64 {
	return v.w
}

func Point(x, y, z float64) Tuple {
	return &PointArray{x, y, z, 1.0, "point"}
}

func Vector(x, y, z float64) Tuple {
	return &VectorArray{x, y, z, 0.0, "vector"}
}

func GetTuple(tupleType string, x, y, z float64) (Tuple, error) {
	if tupleType == "point" {
		return Point(x, y, z), nil
	}
	if tupleType == "vector" {
		return Vector(x, y, z), nil
	}
	return nil, fmt.Errorf("unsupported tuple type: %s", tupleType)
}

func Add(a, b Tuple) (Tuple, error) {
	if a.GetW() == 1 && b.GetW() == 1 {
		return nil, fmt.Errorf("cannot add tuples with w coordinate as 1")
	}

	return &PointArray{
		a.GetX() + b.GetX(),
		a.GetY() + b.GetY(),
		a.GetZ() + b.GetZ(),
		a.GetW() + b.GetW(),
		"point",
	}, nil
}

func Subtract(a, b Tuple) (Tuple, error) {
	if a.GetProperty() == "vector" && b.GetProperty() == "point" {
		return nil, fmt.Errorf("cannot subtract point from vector")
	}

	w := a.GetW() - b.GetW()

	var property string
	if w == 1.0 {
		property = "point"
	} else {
		property = "vector"
	}

	if property == "vector" {
		return &VectorArray{
			a.GetX() - b.GetX(),
			a.GetY() - b.GetY(),
			a.GetZ() - b.GetZ(),
			w,
			property,
		}, nil
	}

	return &PointArray{
		a.GetX() - b.GetX(),
		a.GetY() - b.GetY(),
		a.GetZ() - b.GetZ(),
		w,
		property,
	}, nil
}

func Negate(a Tuple) (Tuple, error) {
	if a.GetProperty() == "point" {
		return &PointArray{
			-a.GetX(),
			-a.GetY(),
			-a.GetZ(),
			-a.GetW(),
			"point",
		}, nil
	}
	return &VectorArray{
		-a.GetX(),
		-a.GetY(),
		-a.GetZ(),
		-a.GetW(),
		"vector",
	}, nil
}

func Multiply(a Tuple, b float64) (Tuple, error) {
	if a.GetProperty() == "point" {
		return &PointArray{
			a.GetX() * b,
			a.GetY() * b,
			a.GetZ() * b,
			a.GetW() * b,
			"point",
		}, nil
	}
	return &VectorArray{
		a.GetX() * b,
		a.GetY() * b,
		a.GetZ() * b,
		a.GetW() * b,
		"vector",
	}, nil
}
