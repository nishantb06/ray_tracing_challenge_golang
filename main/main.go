package main

import (
	"fmt"

	"github.com/nishantb06/ray_tracing_challenge_golang/features"
)

func main() {
	fmt.Println("Hello, World!")

	c := features.Canvas_(224, 224)
	ppm := features.CanvasToPPM(c)
	fmt.Println(ppm)
	// save ppm to file
	features.PPMToFile(ppm, "canvas.ppm")

}
