package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"runtime"
)

type line struct {
	node
	colorful

	vao, vbo, ibo uint32
	indices       int32

	program Program

	thickness float32
}

func NewLine(program Program, thickness float32, vertices []float32, indices []uint32) Primitive {
	l := &line{
		node: node{
			scl: mgl32.Vec3{1., 1., 1.},
		},
		colorful: colorful{
			color: mgl32.Vec4{0., 1., 0., 1.},
		},
		program:   program,
		thickness: thickness,
		indices:   int32(len(indices)),
	}
	l.glLine(vertices, indices)
	return l
}

func (l *line) Update(project, camera mgl32.Mat4) {
	transform := mgl32.Translate3D(l.pos[0], l.pos[1], l.pos[2])
	scale := mgl32.Scale3D(l.scl[0], l.scl[1], l.scl[2])
	rotation := mgl32.AnglesToQuat(l.rot[0], l.rot[1], l.rot[2], mgl32.XYZ)

	transform = transform.Mul4(scale).Mul4(rotation.Mat4())

	l.program.SetUniformValue("thickness", l.thickness)
	l.program.SetUniformValue("view", camera)
	l.program.SetUniformValue("model", transform)
	l.program.SetUniformValue("project", project)
	l.program.SetUniformValue("color", l.color)
}

func (l *line) Render() {
	l.program.Apply()

	gl.BindVertexArray(l.vao)
	gl.DrawElements(gl.TRIANGLES, l.indices, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}

func (l *line) glLine(vertices []float32, indices []uint32) {
	gl.GenVertexArrays(1, &l.vao)
	gl.GenBuffers(1, &l.vbo)
	gl.GenBuffers(1, &l.ibo)

	runtime.SetFinalizer(l, func(l *line) {
		gl.DeleteBuffers(1, &l.ibo)
		gl.DeleteBuffers(1, &l.vbo)
		gl.DeleteVertexArrays(1, &l.vao)
	})

	gl.BindVertexArray(l.vao)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, l.ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, l.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 2 * (3 * 4)
	var offset = 0

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 3 * 4

	// normal
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(1)
	offset += 3 * 4

	gl.BindVertexArray(0)
}
