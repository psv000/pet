package graphics

import "github.com/go-gl/mathgl/mgl32"

type Colorful interface {
	SetColor(color mgl32.Vec4)
}

type colorful struct {
	color mgl32.Vec4
}

func (c *colorful) SetColor(color mgl32.Vec4) {
	c.color = color
}
