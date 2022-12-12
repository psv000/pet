package solid

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"pet/internal/app"
	"pet/internal/graphics"
	"pet/internal/physics"
)

type Object struct {
	graphics.Primitive
	*physics.Object

	// - service -
	velocityVec app.Vector
	forceVec    app.Vector
}

func NewObject(g graphics.Primitive, p *physics.Object) *Object {
	velocity := app.NewVector()
	velocity.SetColor(mgl32.Vec4{0., 165. / 255., 0., 0.7})
	velocity.SetPosition(mgl32.Vec3{0., 0., -1.5})

	force := app.NewVector()
	force.SetColor(mgl32.Vec4{1., 0., 0., 0.5})
	force.SetPosition(mgl32.Vec3{0., 0., -1.})

	return &Object{
		Primitive: g, Object: p,
		velocityVec: velocity,
		forceVec:    force,
	}
}

func (obj *Object) SetPosition(pos mgl32.Vec3) {
	obj.Primitive.SetPosition(pos)
	obj.velocityVec.SetPosition(pos)
	obj.forceVec.SetPosition(pos)
	obj.P = Vec3H(pos)
}

func (obj *Object) Update(dt float64, projectTransform, camTransform mgl32.Mat4) {
	x, y, z := obj.P.Elem()
	pos := mgl32.Vec3{float32(x), float32(y), float32(z)}
	obj.SetPosition(pos)

	obj.Primitive.Update(projectTransform, camTransform)
	obj.Render()

	obj.velocityVec.SetDirection(Vec3L(obj.V.Normalize()))
	obj.velocityVec.SetLength(float32(obj.V.Len()))

	obj.velocityVec.Update(projectTransform, camTransform)
	obj.velocityVec.Render()
}

func Vec3L(vec mgl64.Vec3) mgl32.Vec3 {
	x, y, z := vec.Elem()
	return mgl32.Vec3{float32(x), float32(y), float32(z)}
}

func Vec3H(vec mgl32.Vec3) mgl64.Vec3 {
	x, y, z := vec.Elem()
	return mgl64.Vec3{float64(x), float64(y), float64(z)}
}
