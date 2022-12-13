package graphics

import "github.com/go-gl/mathgl/mgl32"

type Node interface {
	SetPosition(pos mgl32.Vec3)
	SetRotation(rotation mgl32.Vec3)
	SetScale(scale float32)
}

type node struct {
	scl mgl32.Vec3
	pos mgl32.Vec3
	rot mgl32.Vec3
}

func (n *node) SetPosition(pos mgl32.Vec3) {
	n.pos = pos
}

func (n *node) SetScale(scale float32) {
	n.scl = mgl32.Vec3{scale, scale, scale}
}

func (n *node) SetRotation(rotation mgl32.Vec3) {
	n.rot = rotation
}
