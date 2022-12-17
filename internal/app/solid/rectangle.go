package solid

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"pet/internal/graphics"
	"pet/internal/resources"
)

type Rectangle struct {
	graphics.Primitive
}

func NewRectangle(position, size mgl32.Vec3) graphics.Primitive {
	d := &Rectangle{
		Primitive: graphics.NewPrimitive(
			resources.MeshShader,
			[]float32{
				0, 0, 0.,
				size[0], 0, 0.,
				size[0], size[1], 0.,
				0, size[1], 0.,
			},
			[]uint32{0, 1, 2, 0, 2, 3},
			graphics.VertexAttributes{
				Stride: 3 * 4,
				List: []graphics.VertexAttribute{
					{
						Size:   3,
						Type:   gl.FLOAT,
						Offset: 0,
					},
				},
			},
		),
	}
	return d
}
