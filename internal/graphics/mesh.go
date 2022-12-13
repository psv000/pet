package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"runtime"
)

type mesh struct {
	node
	colorful

	vao, vbo, ibo uint32
	indices       int32

	program Program
}

func NewPrimitive(shader Program, vertices []float32, indices []uint32) Primitive {
	p := &mesh{
		node: node{
			scl: mgl32.Vec3{1., 1., 1.},
		},
		program: shader,
		indices: int32(len(indices)),
	}
	p.glMesh(vertices, indices)
	return p
}

func (m *mesh) Update(project, camera mgl32.Mat4) {
	transform := mgl32.Translate3D(m.pos[0], m.pos[1], m.pos[2])
	scale := mgl32.Scale3D(m.scl[0], m.scl[1], m.scl[2])
	rotation := mgl32.AnglesToQuat(m.rot[0], m.rot[1], m.rot[2], mgl32.XYZ)

	transform = transform.Mul4(scale).Mul4(rotation.Mat4())

	m.program.SetUniformValue("view", camera)
	m.program.SetUniformValue("model", transform)
	m.program.SetUniformValue("project", project)
	m.program.SetUniformValue("color", m.color)
}

func (m *mesh) Render() {
	m.program.Apply()

	gl.BindVertexArray(m.vao)
	gl.DrawElements(gl.TRIANGLES, m.indices, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}

func (m *mesh) glMesh(vertices []float32, indices []uint32) {
	gl.GenVertexArrays(1, &m.vao)
	gl.GenBuffers(1, &m.vbo)
	gl.GenBuffers(1, &m.ibo)

	runtime.SetFinalizer(m, func(m *mesh) {
		gl.DeleteBuffers(1, &m.ibo)
		gl.DeleteBuffers(1, &m.vbo)
		gl.DeleteVertexArrays(1, &m.vao)
	})

	gl.BindVertexArray(m.vao)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 3*4 + 2*4
	var offset = 0

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 3 * 4

	// uv position
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(1)
	offset += 2 * 4

	gl.BindVertexArray(0)
}
