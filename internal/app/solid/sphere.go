package solid

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"pet/internal/graphics"
	"pet/internal/physics/collisions"
)

type Sphere struct {
	graphics graphics.Primitive
	physics  collisions.Sphere

	// - service -
	velocity Line
}

func NewSphere(g graphics.Primitive, p collisions.Sphere) *Sphere {
	return &Sphere{
		graphics: g, physics: p,
		velocity: NewLine(),
	}
}

func (obj *Sphere) Physics() collisions.Sphere {
	return obj.physics
}

func (obj *Sphere) Graphics() graphics.Primitive {
	return obj.graphics
}

func (obj *Sphere) SetPosition(pos mgl32.Vec3) {
	obj.graphics.SetPosition(pos)
	obj.velocity.SetPosition(pos)
	obj.physics.SetPosition(Vec3H(pos))
}

func (obj *Sphere) Update(dt float64, projectTransform, camTransform mgl32.Mat4) {
	x, y, z := obj.physics.Position().Elem()
	pos := mgl32.Vec3{float32(x), float32(y), float32(z)}

	{
		obj.velocity.SetLength(float32(obj.physics.Velocity().Len()))
	}
	{
		d := mgl64.Vec3{1., 0., 0.}.Dot(obj.physics.Velocity())
		cos := d / obj.physics.Velocity().Len()
		diff := math.Abs(obj.physics.Velocity().Y())
		if diff == 0 {
			diff = 1
		}
		angle := math.Acos(cos) * obj.physics.Velocity().Y() / diff
		obj.velocity.SetRotation(mgl32.Vec3{0., 0., float32(angle)})
	}

	obj.SetPosition(pos)

	obj.graphics.Update(projectTransform, camTransform)
	obj.graphics.Render()

	obj.velocity.Update(projectTransform, camTransform)
	obj.velocity.Render()
}

func Vec3L(vec mgl64.Vec3) mgl32.Vec3 {
	x, y, z := vec.Elem()
	return mgl32.Vec3{float32(x), float32(y), float32(z)}
}

func Vec3H(vec mgl32.Vec3) mgl64.Vec3 {
	x, y, z := vec.Elem()
	return mgl64.Vec3{float64(x), float64(y), float64(z)}
}