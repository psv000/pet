package app

import (
	"github.com/go-gl/mathgl/mgl32"
	"pet/internal/graphics"
	"pet/internal/resources"
)

type Vector interface {
	graphics.Primitive

	SetLength(l float32)
	SetDirection(d mgl32.Vec3)
}

type vector struct {
	graphics.Line

	length    float32
	direction mgl32.Vec3
}

func NewVector() Vector {
	v := &vector{
		Line: graphics.NewLine(
			resources.VectorShader,
			[]float32{
				0., 0., 0.,
				1., 0., 0.,
			},
			[]uint32{0, 1}),
	}
	v.SetColor(mgl32.Vec4{1., 1., 0., 1.})
	return v
}

func (v *vector) SetLength(l float32) {
	v.length = l
}

func (v *vector) SetDirection(d mgl32.Vec3) {
	v.direction = d
}
