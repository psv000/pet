package assets

import _ "embed"

// - Shaders -

//go:embed shaders/general.vert
var BasicVertexShader string

//go:embed shaders/general.frag
var BasicFragmentShader string

//go:embed shaders/dot.vert
var DotVertexShader string

//go:embed shaders/dot.frag
var DotFragmentShader string

//go:embed shaders/line.vert
var LineVertexShader string

// - Fonts -

//go:embed fonts/luxisr.ttf
var DefaultFont string
