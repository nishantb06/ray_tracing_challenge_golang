package main

import (
	"fmt"

	"github.com/nishantb06/ray_tracing_challenge_golang/features"
)

func main() {
	fmt.Println("Hello, World!")

	// c := features.Canvas_(224, 224)
	// ppm := features.CanvasToPPM(c)
	// fmt.Println(ppm)
	// save ppm to file
	c := features.Canvas_(10, 2)
	c1 := features.Color_(1, 0.8, 0.6)
	for i := range c.Pixels {
		for j := range c.Pixels[i] {
			c.WritePixel(j, i, c1)
		}
	}
	ppm := features.CanvasToPPM(c)
	features.PPMToFile(ppm, "canvas.ppm")

}
