package resources

import (
	"pet/assets"
	"pet/internal/graphics"
)

// - shaders -
var (
	DotShader  graphics.Program
	LineShader graphics.Program
)

func Load() {
	loadShaders()
}

func loadShaders() {
	DotShader = graphics.NewProgram(
		assets.BasicVertexShader,
		assets.DotFragmentShader,
		[]string{
			"model", "view", "project",
			"color",
		},
	)

	//LineShader = graphics.NewProgram(
	//	assets.LineVertexShader,
	//	assets.AlbedoFragmentShader,
	//	[]string{
	//		"model", "view", "project",
	//		"color",
	//	},
	//)
}
