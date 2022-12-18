package collisions

import (
	"math"

	"pet/internal/physics/model"

	"github.com/go-gl/mathgl/mgl64"
)

type Sphere interface {
	Model

	SetRadius(radius float64)
	Radius() float64
}

type sphere struct {
	model.Core
	radius float64
}

func NewSphere() Sphere {
	return &sphere{
		radius: 1.,
	}
}

func (obj *sphere) SetRadius(radius float64) {
	obj.radius = radius
}

func (obj *sphere) Radius() float64 {
	return obj.radius
}

func DetectSpheresCollision(obj1, obj2 Sphere) (bool, float64) {
	l := obj2.Position().Sub(obj1.Position())
	dsq := l.Dot(l)
	r := obj1.Radius() + obj2.Radius()
	rsq := r * r
	return dsq < rsq, r - math.Sqrt(dsq)
}

func AmendSpheres(penetration float64, obj1, obj2 Sphere) {
	normal := obj1.Position().Sub(obj2.Position()).Normalize()
	correction := normal.Mul(penetration / (obj1.Mass() + obj2.Mass()))

	obj1.SetPosition(obj1.Position().Add(correction.Mul(obj1.Mass())))
	obj2.SetPosition(obj2.Position().Sub(correction.Mul(obj2.Mass())))
}

func ResolveSpheresCollision(dt float64, obj1, obj2 Sphere) {
	vec := obj1.Position().Sub(obj2.Position())

	n := vec.Normalize()
	tan := mgl64.Vec3{-n[1], n[0], 0}

	dpTan1 := obj1.Velocity()[0]*tan[0] + obj1.Velocity()[1]*tan[1]
	dpTan2 := obj2.Velocity()[0]*tan[0] + obj2.Velocity()[1]*tan[1]

	dpNorm1 := obj1.Velocity()[0]*n[0] + obj1.Velocity()[1]*n[1]
	dpNorm2 := obj2.Velocity()[0]*n[0] + obj2.Velocity()[1]*n[1]

	m1 := (dpNorm1*(obj1.Mass()-obj2.Mass()) + 2.*obj2.Mass()*dpNorm2) / (obj1.Mass() + obj2.Mass())
	m2 := (dpNorm2*(obj2.Mass()-obj1.Mass()) + 2.*obj1.Mass()*dpNorm1) / (obj1.Mass() + obj2.Mass())

	e := obj2.Restitution()
	if obj1.Restitution() < obj1.Restitution() {
		e = obj1.Restitution()
	}

	obj1Vel := tan.Mul(dpTan1).Add(n.Mul(m1 * e))
	obj2Vel := tan.Mul(dpTan2).Add(n.Mul(m2 * e))

	obj1.SetVelocity(obj1Vel)
	obj2.SetVelocity(obj2Vel)
}

func DetectSphereVsCuboidCollision(c Sphere, p CuboidCollided) (bool, float64, mgl64.Vec3) {
	pc := p.Position().Add(p.Size().Mul(0.5))
	n := c.Position().Sub(pc)

	closest := n

	extent := p.Size().Mul(0.5)
	closest = clampVec3(closest, extent.Mul(-1), extent)

	var inside bool
	if n.ApproxEqualThreshold(closest, 1e-9) {
		inside = true

		if mgl64.Abs(n[0]) > mgl64.Abs(n[1]) {
			if closest[0] > 0 {
				closest[0] = extent[0]
			} else {
				closest[0] = -extent[0]
			}
		} else {
			if closest[1] > 0 {
				closest[1] = extent[1]
			} else {
				closest[1] = -extent[1]
			}
		}
	}

	normal := n.Sub(closest)
	d := normal.Dot(normal)
	rsq := c.Radius()
	rsq *= rsq

	if d > rsq && !inside {
		return false, 0., mgl64.Vec3{}
	}

	closest = p.Position().Add(p.Size().Mul(0.5)).Add(closest)
	return true, c.Radius() - closest.Sub(c.Position()).Len(), closest
}

func ResolveSphereVsCuboidCollision(c Sphere, p CuboidCollided, point mgl64.Vec3) {
	rv := c.Velocity().Sub(p.Velocity())
	n := c.Position().Sub(point).Normalize()
	velProjection := rv.Dot(n)

	if velProjection > 0 {
		return
	}

	e := p.Restitution()
	if c.Restitution() < p.Restitution() {
		e = c.Restitution()
	}

	j := (1 + e) * velProjection
	j /= 1/c.Mass() + 1/p.Mass()

	impulse := n.Mul(j)

	c.AddVelocity(impulse.Mul(-1. / c.Mass()))
}

func clampVec3(a, b, c mgl64.Vec3) (result mgl64.Vec3) {
	result[0] = mgl64.Clamp(a[0], b[0], c[0])
	result[1] = mgl64.Clamp(a[1], b[1], c[1])
	result[2] = mgl64.Clamp(a[2], b[2], c[2])
	return
}
