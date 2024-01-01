package main

import (
	"fmt"

	"github.com/nishantb06/ray_tracing_challenge_golang/features"
	"github.com/nishantb06/ray_tracing_challenge_golang/utils"
)

func sample_conversion() {
	// Specify the paths for the input PPM file and the output PNG file
	ppmPath := "projectile.ppm"
	pngPath := "projectile.png"

	// Convert PPM to PNG
	if err := utils.ConvertPPMtoPNG(ppmPath, pngPath); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Conversion successful!")
}

// func simulateProjectile(start features.Tuple, velocity features.Tuple, multiplier float64, wind features.Tuple, gravity features.Tuple, canvasWidth, canvasHeight int, filename string) {
// 	velocity, _ = features.Normalize(velocity)
// 	velocity, _ = features.Multiply(velocity, multiplier)
// 	p := features.Projectile{Position: start, Velocity: velocity}
// 	e := features.Environment{Gravity: gravity, Wind: wind}

// 	c := features.Canvas_(canvasWidth, canvasHeight)
// 	for p.Position.GetY() > 0 {
// 		c = features.PlotProjectile(c, p)
// 		p = features.Tick(e, p)
// 	}
// 	ppm := features.CanvasToPPM(c)
// 	features.PPMToFile(ppm, filename)
// }

func main() {
	fmt.Println("Hello, World!")

	start := features.Point(0, 1, 0)
	velocity := features.Vector(1, 1.8, 0)
	gravity := features.Vector(0, -0.1, 0)
	wind := features.Vector(-0.01, 0, 0)

	features.SimulateProjectile(start, velocity, 11.25, wind, gravity, 900, 550, "projectile3.ppm")
}
