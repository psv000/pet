package resources

import (
	"pet/assets"
	"pet/internal/graphics"
)

// - shaders -
var (
	DotShader  graphics.Program
	LineShader graphics.Program
	MeshShader graphics.Program

	ProgramList []graphics.Program
)

func Load() {
	loadPrograms()
}

func newProgram(vertex, fragment string, uniforms []string) graphics.Program {
	program := graphics.NewProgram(
		vertex,
		fragment,
		uniforms,
	)
	ProgramList = append(ProgramList, program)
	return program
}

func loadPrograms() {
	DotShader = newProgram(
		assets.DotVertexShader,
		assets.DotFragmentShader,
		[]string{
			"model", "view", "project",
			"color",
		},
	)

	LineShader = newProgram(
		assets.LineVertexShader,
		assets.BasicFragmentShader,
		[]string{
			"thickness", "length", "model", "view", "project",
			"color",
		},
	)

	MeshShader = newProgram(
		assets.BasicVertexShader,
		assets.BasicFragmentShader,
		[]string{
			"model", "view", "project",
			"color",
		},
	)
}
