package collisions

import (
	"pet/internal/physics/model"

	"github.com/go-gl/mathgl/mgl64"
)

type CuboidCollided interface {
	Model

	Size() mgl64.Vec3
	SetSize(mgl64.Vec3)
}

type cuboid struct {
	model.Core

	size mgl64.Vec3
}

func NewCuboid() CuboidCollided {
	return &cuboid{}
}

func (c *cuboid) Size() mgl64.Vec3 {
	return c.size
}

func (c *cuboid) SetSize(s mgl64.Vec3) {
	c.size = s
}
