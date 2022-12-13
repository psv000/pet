package solid

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"pet/internal/graphics"
	"pet/internal/physics"
)

type Object struct {
	graphics.Primitive
	*physics.Object

	// - service -
	velocity Line
}

func NewObject(g graphics.Primitive, p *physics.Object) *Object {
	return &Object{
		Primitive: g, Object: p,
		velocity: NewLine(),
	}
}

func (obj *Object) SetPosition(pos mgl32.Vec3) {
	obj.Primitive.SetPosition(pos)
	obj.velocity.SetPosition(pos)
	obj.P = Vec3H(pos)
}

func (obj *Object) Update(dt float64, projectTransform, camTransform mgl32.Mat4) {
	x, y, z := obj.P.Elem()
	pos := mgl32.Vec3{float32(x), float32(y), float32(z)}
	obj.SetPosition(pos)

	obj.Primitive.Update(projectTransform, camTransform)
	obj.Render()

	{
		d := mgl64.Vec3{1., 0., 0.}.Dot(obj.V)
		cos := d / obj.V.Len()
		angle := math.Acos(cos) * obj.V.Y() / math.Abs(obj.V.Y())
		obj.velocity.SetRotation(mgl32.Vec3{0., 0., float32(angle)})
	}
	{
		len := float32(obj.V.Len())
		color1, color2 := mgl32.Vec4{0., 1., 0., 1.}, mgl32.Vec4{1., 0., 0., 1.}
		color := color2.Sub(color1).Mul(len * 2).Add(color1)
		obj.velocity.SetColor(color)
	}

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
