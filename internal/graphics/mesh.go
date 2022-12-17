package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"runtime"
	"unsafe"
)

type mesh struct {
	node
	colorful
	programmable

	vao, vbo, ibo uint32
	indices       int32
}

func NewPrimitive(program Program, vertices []float32, indices []uint32, attributes VertexAttributes) Primitive {
	p := &mesh{
		node: node{
			scl: mgl32.Vec3{1., 1., 1.},
		},
		programmable: programmable{
			program: program,
		},
		indices: int32(len(indices)),
	}
	p.glMesh(vertices, indices, attributes)
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

func (m *mesh) Program() Program {
	return m.program
}

func (m *mesh) glMesh(vertices []float32, indices []uint32, attributes VertexAttributes) {
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

	for i, attr := range attributes.List {
		gl.VertexAttribPointer(uint32(i), attr.Size, attr.Type, attr.Normalized, attributes.Stride, unsafe.Pointer(attr.Offset))
		gl.EnableVertexAttribArray(uint32(i))
	}

	gl.BindVertexArray(0)
}
