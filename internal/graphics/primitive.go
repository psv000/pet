package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Node interface {
	SetPosition(pos mgl32.Vec3)
	SetScale(scale float32)
}

type Primitive interface {
	Node

	Render()
	Update(project, camera mgl32.Mat4)

	SetColor(color mgl32.Vec4)
}

type primitive struct {
	vao     uint32
	indices int32

	shader Shader

	scl mgl32.Vec3
	pos mgl32.Vec3
	rot mgl32.Vec3

	color mgl32.Vec4
}

func NewPrimitive(shader Shader, vertices []float32, indices []uint32) Primitive {
	p := &primitive{
		shader:  shader,
		indices: int32(len(indices)),
	}
	p.glRegister(vertices, indices)
	return p
}

func (p *primitive) SetPosition(pos mgl32.Vec3) {
	p.pos = pos
}

func (p *primitive) SetScale(scale float32) {
	p.scl = mgl32.Vec3{scale, scale, scale}
}

func (p *primitive) SetColor(color mgl32.Vec4) {
	p.color = color
}

func (p *primitive) Update(project, camera mgl32.Mat4) {
	transform := mgl32.Translate3D(p.pos[0], p.pos[1], p.pos[2])
	scale := mgl32.Scale3D(p.scl[0], p.scl[1], p.scl[2])
	rotation := mgl32.AnglesToQuat(p.rot[0], p.rot[1], p.rot[2], mgl32.XYZ)

	transform = transform.Mul4(scale).Mul4(rotation.Mat4())

	p.shader.SetUniformValue("view", camera)
	p.shader.SetUniformValue("model", transform)
	p.shader.SetUniformValue("project", project)
	p.shader.SetUniformValue("color", p.color)
}

func (p *primitive) Render() {
	p.shader.Apply()

	gl.BindVertexArray(p.vao)
	gl.DrawElements(gl.TRIANGLES, p.indices, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}

func (p *primitive) Shader() Shader {
	return p.shader
}

func (p *primitive) glRegister(vertices []float32, indices []uint32) {
	gl.GenVertexArrays(1, &p.vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)

	var ibo uint32
	gl.GenBuffers(1, &ibo)

	gl.BindVertexArray(p.vao)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	var stride int32 = 3*4 + 2*4
	var offset int = 0

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
