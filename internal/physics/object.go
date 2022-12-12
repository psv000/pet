package physics

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Object struct {
	M float64
	V mgl64.Vec3
	P mgl64.Vec3
	R float64
}

func (obj *Object) Position() mgl64.Vec3 {
	return obj.P
}

func (obj *Object) Velocity() mgl64.Vec3 {
	return obj.V
}

func (obj *Object) Mass() float64 {
	return obj.M
}

// Restitution is a упругость
func (obj *Object) Restitution() float64 {
	return obj.M * 1.6
}

func (obj *Object) Radius() float64 {
	return obj.R
}

func (obj *Object) SetVelocity(vel mgl64.Vec3) {
	obj.V = vel
}

func (obj *Object) ApplyForce(dt float64, force mgl64.Vec3) {
	a := force.Mul(1. / obj.M)
	obj.V = obj.V.Add(a.Mul(dt))
}

func (obj *Object) Update(dt float64) {
	s := obj.V.Mul(dt)
	obj.P = obj.P.Add(s)
}
