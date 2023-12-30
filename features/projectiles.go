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
