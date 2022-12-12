package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader interface {
	ID() uint32
	Apply()
	SetUniformValue(name string, value any)
}

type shader struct {
	descriptor uint32
	uniforms   map[string]int32
	values     map[string]any
}

func NewShader(vertex, fragment string, uniforms []string) Shader {
	s := &shader{
		uniforms: make(map[string]int32, 16),
		values:   make(map[string]any, 16),
	}
	s.init(vertex, fragment, uniforms)
	return s
}

func (s *shader) init(vertex, fragment string, uniforms []string) {
	shaders := compileShaders(vertex, fragment)
	s.descriptor = linkShaders(shaders)

	gl.UseProgram(s.descriptor)
	for _, name := range uniforms {
		s.uniforms[name] = gl.GetUniformLocation(s.descriptor, gl.Str(name+"\x00"))
	}
}

func (s *shader) ID() uint32 {
	return s.descriptor
}

func (s *shader) Apply() {
	gl.UseProgram(s.descriptor)
	for name := range s.uniforms {
		loc := s.uniforms[name]
		if val, found := s.values[name]; found {
			switch t := val.(type) {
			case mgl32.Mat4:
				gl.UniformMatrix4fv(loc, 1, false, &t[0])
			case mgl32.Vec4:
				gl.Uniform4f(loc, t[0], t[1], t[2], t[3])
			case mgl32.Vec3:
				gl.Uniform3f(loc, t[0], t[1], t[2])
			case float32:
				gl.Uniform1f(loc, t)
			}
		}
	}
}

func (s *shader) SetUniformValue(name string, value any) {
	if _, found := s.uniforms[name]; found {
		s.values[name] = value
	}
}
