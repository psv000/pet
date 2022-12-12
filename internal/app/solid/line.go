package solid

import (
	"github.com/go-gl/mathgl/mgl32"
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
			1.,
			[]float32{
				0., 0., 0.,
				1., 0., 0.,
			}),
	}
	l.SetColor(mgl32.Vec4{1., 1., 0., 1.})
	return l
}
