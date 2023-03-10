package solid

import (
	"github.com/go-gl/gl/v4.1-core/gl"
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
			[]uint32{0, 1, 2, 0, 2, 3},
			graphics.VertexAttributes{
				Stride: 3*4 + 2*4,
				List: []graphics.VertexAttribute{
					{
						Size:   3,
						Type:   gl.FLOAT,
						Offset: 0,
					},
					{
						Size:   2,
						Type:   gl.FLOAT,
						Offset: 3 * 4,
					},
				},
			},
		),
	}
	return d
}
