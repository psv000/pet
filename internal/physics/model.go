package physics

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Model interface {
	Position() mgl64.Vec3
	Rotation() mgl64.Vec3
	Velocity() mgl64.Vec3

	Mass() float64
	Restitution() float64

	SetPosition(position mgl64.Vec3)
	SetVelocity(vel mgl64.Vec3)

	SetMass(mass float64)
	SetRestitution(restitution float64)

	AddVelocity(vec3 mgl64.Vec3)

	ApplyForce(dt float64, force mgl64.Vec3)
	Update(dt float64)
}

type model struct {
	mass, restitution  float64
	velocity, force    mgl64.Vec3
	position, rotation mgl64.Vec3
}

func (m *model) Position() mgl64.Vec3 {
	return m.position
}

func (m *model) Rotation() mgl64.Vec3 {
	return m.rotation
}

func (m *model) Velocity() mgl64.Vec3 {
	return m.velocity
}

func (m *model) Mass() float64 {
	return m.mass
}

func (m *model) Restitution() float64 {
	return m.restitution
}

func (m *model) SetPosition(p mgl64.Vec3) {
	m.position = p
}

func (m *model) SetVelocity(v mgl64.Vec3) {
	m.velocity = v
}

func (m *model) SetMass(mass float64) {
	m.mass = mass
}

func (m *model) SetRestitution(restitution float64) {
	m.restitution = restitution
}

func (m *model) AddVelocity(v mgl64.Vec3) {
	m.velocity = m.velocity.Add(v)
}

func (m *model) ApplyForce(dt float64, force mgl64.Vec3) {
	m.force = force
	a := m.force.Mul(1. / m.mass)
	m.velocity = m.velocity.Add(a.Mul(dt))
}

func (m *model) Update(dt float64) {
	s := m.velocity.Mul(dt)
	m.position = m.position.Add(s)
}
