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
}
