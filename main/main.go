package main

import (
	"fmt"
	"math"

	"github.com/nishantb06/ray_tracing_challenge_golang/features"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Abs(-1))
	p, error := features.GetTuple("point", 1, 2, 3)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(p.GetX(), p.GetY(), p.GetZ(), p.GetW(), p.GetProperty())
	fmt.Println(features.GetTuple("point", 2, 2, 3))

	normalizedVector, _ := features.Normalize(features.Vector(1, 1, 0))
	normalizedVector, _ = features.Multiply(normalizedVector, 0.2)
	proj := features.Projectile{Position: features.Point(0, 1, 0), Velocity: normalizedVector}
	env := features.Environment{Gravity: features.Vector(0, -0.1, 0), Wind: features.Vector(-0.01, 0, 0)}
	fmt.Println(proj, env)

	for proj.Position.GetY() > 0 {
		proj = features.Tick(env, proj)
		fmt.Println(proj.Position.GetX(), proj.Position.GetY(), proj.Position.GetZ())
	}
	multiLineString :=
		`This is a
multi-line
string in Go.
It preserves
line breaks.`
	line2 := "This is a\nmulti-line\nstring in Go.\nIt preserves\nline breaks."
	fmt.Println(multiLineString)
	fmt.Println(line2)
	fmt.Println("==================================================")
	fmt.Println(multiLineString + "\n" + line2)
}
