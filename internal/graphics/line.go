package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Line interface {
	Primitive
}

type line struct {
	primitive
}

func NewLine(shader Shader, vertices []float32, indices []uint32) Line {
	l := &line{
		primitive: primitive{
			shader:  shader,
			indices: int32(len(indices)),
		},
	}
	l.glRegister(vertices, indices)
	return l
}

func (l *line) Render() {
	l.shader.Apply()

	gl.BindVertexArray(l.vao)
	gl.DrawElements(gl.LINE_STRIP, l.indices, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}
