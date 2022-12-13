package collisions

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

type SphereCollided interface {
	Collided
	Radius() float64
}

func Spheres(c1, c2 SphereCollided) (bool, float64) {
	l := c2.Position().Sub(c1.Position())
	dsq := l.Dot(l)
	r := c1.Radius() + c2.Radius()
	rsq := r * r
	return dsq <= rsq, r - math.Sqrt(dsq)
}

func ResolveSpheres(c1, c2 Collided, dt float64) (ƒ1 mgl64.Vec3, ƒ2 mgl64.Vec3) {
	rv := c1.Velocity().Sub(c2.Velocity())
	p := c1.Position().Sub(c2.Position())
	n := p.Normalize()

	velocityAlongNormal := rv.Dot(n)

	if velocityAlongNormal > 0 {
		return
	}

	var e float64
	if c1.Restitution() < c2.Restitution() {
		e = c1.Restitution()
	} else {
		e = c2.Restitution()
	}

	j := (1 + e) * velocityAlongNormal
	j /= 1./c1.Mass() + 1./c2.Mass()

	impulse := n.Mul(j)

	factor := 1.
	return impulse.Mul(-1. / c1.Mass() * factor).Mul(1. / dt),
		impulse.Mul(1. / c2.Mass() * factor).Mul(1. / dt)

}
