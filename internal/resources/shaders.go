package resources

import (
	"pet/assets"
	"pet/internal/graphics"
)

// - shaders -
var (
	DotShader    graphics.Shader
	VectorShader graphics.Shader
)

func Load() {
	loadShaders()
}

func loadShaders() {
	DotShader = graphics.NewShader(
		assets.BasicVertexShader,
		assets.DotFragmentShader,
		[]string{
			"model", "view", "project",
			"color",
		},
	)

	VectorShader = graphics.NewShader(
		assets.BasicVertexShader,
		assets.AlbedoFragmentShader,
		[]string{
			"model", "view", "project",
			"color",
		},
	)
}
