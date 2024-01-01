package features

type Projectile struct {
	Position Tuple
	Velocity Tuple
}

type Environment struct {
	Gravity Tuple
	Wind    Tuple
}

func Tick(env Environment, proj Projectile) Projectile {
	position, _ := Add(proj.Position, proj.Velocity)
	velocity, _ := Add(proj.Velocity, env.Gravity)
	velocity, _ = Add(velocity, env.Wind)
	return Projectile{position, velocity}
}

func PlotProjectile(canvas Canvas, proj Projectile) Canvas {
	x := int(proj.Position.GetX())
	y := canvas.Height - int(proj.Position.GetY())
	canvas.Pixels[y][x] = Color_(0, 1, 0)
	return canvas
}

func SimulateProjectile(start Tuple, velocity Tuple, multiplier float64, wind Tuple, gravity Tuple, canvasWidth, canvasHeight int, filename string) {
	velocity, _ = Normalize(velocity)
	velocity, _ = Multiply(velocity, multiplier)
	p := Projectile{Position: start, Velocity: velocity}
	e := Environment{Gravity: gravity, Wind: wind}

	c := Canvas_(canvasWidth, canvasHeight)
	for p.Position.GetY() > 0 {
		c = PlotProjectile(c, p)
		p = Tick(e, p)
	}
	ppm := CanvasToPPM(c)
	PPMToFile(ppm, filename)
}
