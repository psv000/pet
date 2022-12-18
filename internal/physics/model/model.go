package model

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Model interface {
	Position() mgl64.Vec3
	Rotation() mgl64.Vec3
	Velocity() mgl64.Vec3

	SetPosition(position mgl64.Vec3)
	SetRotation(rotation mgl64.Vec3)
	SetVelocity(vel mgl64.Vec3)

	Mass() float64
	Restitution() float64

	SetMass(mass float64)
	SetRestitution(restitution float64)

	AddVelocity(vec3 mgl64.Vec3)

	ApplyForce(dt float64, force mgl64.Vec3)
	Update(dt float64)
}

type Core struct {
	mass, restitution  float64
	velocity, force    mgl64.Vec3
	position, rotation mgl64.Vec3
}

func (m *Core) Position() mgl64.Vec3 {
	return m.position
}

func (m *Core) Rotation() mgl64.Vec3 {
	return m.rotation
}

func (m *Core) Restitution() float64 {
	return m.restitution
}

func (m *Core) SetPosition(p mgl64.Vec3) {
	m.position = p
}

func (m *Core) SetRotation(r mgl64.Vec3) {
	m.rotation = r
}

func (m *Core) Velocity() mgl64.Vec3 {
	return m.velocity
}

func (m *Core) Mass() float64 {
	return m.mass
}

func (m *Core) SetVelocity(v mgl64.Vec3) {
	m.velocity = v
}

func (m *Core) SetMass(mass float64) {
	m.mass = mass
}

func (m *Core) SetRestitution(restitution float64) {
	m.restitution = restitution
}

func (m *Core) AddVelocity(v mgl64.Vec3) {
	m.velocity = m.velocity.Add(v)
}

func (m *Core) ApplyForce(dt float64, force mgl64.Vec3) {
	m.force = force
	a := m.force.Mul(1. / m.mass)
	m.velocity = m.velocity.Add(a.Mul(dt))
}

func (m *Core) Update(dt float64) {
	s := m.velocity.Mul(dt)
	m.position = m.position.Add(s)
}
