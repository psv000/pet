package app

import (
	"pet/internal/graphics"
	"pet/internal/resources"
)

type dot struct {
	graphics.Primitive
}

func NewDot() graphics.Primitive {
	d := &dot{
		Primitive: graphics.NewPrimitive(
			resources.DotShader,
			[]float32{
				-0.1, -0.1, 0., -1., -1.,
				+0.1, -0.1, 0., +1., -1.,
				+0.1, +0.1, 0., +1., +1.,
				-0.1, +0.1, 0., -1., +1.,
			},
			[]uint32{0, 1, 2, 0, 2, 3}),
	}
	return d
}
