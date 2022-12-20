package solid

import (
	"github.com/go-gl/gl/v4.1-core/gl"
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

	vel mgl64.Vec3
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
		obj.velocity.SetLength(float32(math.Log10(obj.physics.Velocity().Len() * 10.)))
	}
	{
		obj.vel = obj.physics.Velocity().Sub(obj.vel).Mul(dt * 30).Add(obj.vel)

		d := mgl64.Vec3{1., 0., 0.}.Dot(obj.vel)
		cos := d / obj.vel.Len()
		diff := math.Abs(obj.vel.Y())
		if diff == 0 {
			diff = 1
		}
		angle := math.Acos(cos) * obj.vel.Y() / diff
		obj.velocity.SetRotation(mgl32.Vec3{0., 0., float32(angle)})
	}

	obj.SetPosition(pos)

	obj.graphics.Update(projectTransform, camTransform)
	obj.graphics.Render(gl.TRIANGLES)

	//obj.velocity.Update(projectTransform, camTransform)
	//obj.velocity.Render()
}

func Vec3L(vec mgl64.Vec3) mgl32.Vec3 {
	x, y, z := vec.Elem()
	return mgl32.Vec3{float32(x), float32(y), float32(z)}
}

func Vec3H(vec mgl32.Vec3) mgl64.Vec3 {
	x, y, z := vec.Elem()
	return mgl64.Vec3{float64(x), float64(y), float64(z)}
}
