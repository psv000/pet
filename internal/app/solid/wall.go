package solid

import (
	"github.com/go-gl/mathgl/mgl32"
	"pet/internal/graphics"
	"pet/internal/physics"
)

type Wall struct {
	graphics graphics.Primitive
	physics  physics.CuboidCollided
}

func NewWall(graphics graphics.Primitive, physics physics.CuboidCollided) *Wall {
	return &Wall{
		graphics: graphics,
		physics:  physics,
	}
}

func (obj *Wall) Physics() physics.CuboidCollided {
	return obj.physics
}

func (obj *Wall) Graphics() graphics.Primitive {
	return obj.graphics
}

func (obj *Wall) SetPosition(pos mgl32.Vec3) {
	obj.graphics.SetPosition(pos)
	obj.physics.SetPosition(Vec3H(pos))
}

func (obj *Wall) SetSize(size mgl32.Vec3) {
	obj.physics.SetSize(Vec3H(size))
}

func (obj *Wall) Update(dt float64, projectTransform, camTransform mgl32.Mat4) {
	x, y, z := obj.physics.Position().Elem()
	pos := mgl32.Vec3{float32(x), float32(y), float32(z)}

	obj.SetPosition(pos)

	obj.graphics.Update(projectTransform, camTransform)
	obj.graphics.Render()
}
