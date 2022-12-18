package gravity

import (
	"github.com/go-gl/mathgl/mgl64"
	"pet/internal/physics/model"
)

func Apply(g, dt float64, obj1, obj2 model.Model) {
	dir := obj2.Position().Sub(obj1.Position())
	rsq := dir.Len()
	rsq *= rsq

	ƒ := g * obj1.Mass() * obj2.Mass() / rsq

	ƒ1 := dir.Normalize().Mul(ƒ)
	ƒ2 := ƒ1.Mul(-1.)

	obj1.ApplyForce(dt, ƒ1)
	obj2.ApplyForce(dt, ƒ2)
}

func Global(a, dt float64, obj model.Model) {
	obj.ApplyForce(dt, mgl64.Vec3{0., -1., 0.}.Mul(obj.Mass()*a))
}
