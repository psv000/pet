package collisions

import "github.com/go-gl/mathgl/mgl64"

type Collided interface {
	Position() mgl64.Vec3
	Velocity() mgl64.Vec3
	Mass() float64
	Restitution() float64

	SetVelocity(vec3 mgl64.Vec3)
	ApplyForce(dt float64, vec3 mgl64.Vec3)
}
