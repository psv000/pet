package solid

import (
	"pet/internal/graphics"
	"pet/internal/resources"
)

type Line interface {
	graphics.Primitive
}

type line struct {
	Line
}

func NewLine() Line {
	l := line{
		Line: graphics.NewLine(
			resources.LineShader,
			2.,
			[]float32{
				0., 0., 0., 0., -1., 0.,
				0.1, 0., 0., 0., -1., 0.,
				0.1, 0., 0., 0., 1., 0.,
				0., 0., 0., 0., 1., 0.,
			}, []uint32{
				0, 1, 2, 0, 2, 3,
			}),
	}
	return l
}
