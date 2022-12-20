package graphics

import "github.com/go-gl/mathgl/mgl32"

type Primitive interface {
	Node
	Colorful

	Render(mode int)
	Update(project, camera mgl32.Mat4)

	Program() Program

	Clear()
}
