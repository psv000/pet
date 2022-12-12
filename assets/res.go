package assets

import _ "embed"

// - Shaders -

//go:embed shaders/general.vert
var BasicVertexShader string

//go:embed shaders/albedo.frag
var AlbedoFragmentShader string

//go:embed shaders/dot.frag
var DotFragmentShader string

//go:embed shaders/line.vert
var LineVertexShader string

// - Fonts -

//go:embed fonts/luxisr.ttf
var DefaultFont string
