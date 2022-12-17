package solid

import (
	"pet/internal/graphics"
	"pet/internal/resources"
)

type Line interface {
	graphics.Primitive
	SetLength(length float32)
}

type line struct {
	graphics.Primitive
	length float32
}

func NewLine() Line {
	l := &line{
		Primitive: graphics.NewLine(
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

func (l *line) SetLength(length float32) {
	l.Program().SetUniformValue("length", 2*length)
}
