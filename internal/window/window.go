package window

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	width  int
	height int
	glfw   *glfw.Window

	inputManager  *InputManager
	dt            float64
	lastFrameTime float64
}

func init() {
	// GLFW event handling must be run on the main OS thread
	runtime.LockOSThread()
}

func (w *Window) InputManager() *InputManager {
	return w.inputManager
}

func NewWindow(width, height int, title string) *Window {
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	context, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	context.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	im := NewInputManager()

	context.SetKeyCallback(im.keyCallback)
	context.SetCursorPosCallback(im.mouseCallback)
	context.SetMouseButtonCallback(im.mouseButtonCallback)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	//gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)

	return &Window{
		width:        width,
		height:       height,
		glfw:         context,
		inputManager: im,
	}
}

func (w *Window) Width() int {
	return w.width
}

func (w *Window) Height() int {
	return w.height
}

func (w *Window) ShouldClose() bool {
	return w.glfw.ShouldClose()
}

func (w *Window) Render(draw func(dt float64)) {
	if w.inputManager.IsActive(ProgramQuit) {
		w.glfw.SetShouldClose(true)
	}

	now := glfw.GetTime()
	w.dt = now - w.lastFrameTime
	w.lastFrameTime = now

	gl.ClearColor(0., 0., 0., 1.)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	draw(w.dt)

	w.inputManager.CheckpointCursorChange()

	w.glfw.SwapBuffers()
}

func (w *Window) BeforeRender() {
	glfw.PollEvents()
}

func (w *Window) ResetFrameTime() {
	w.lastFrameTime = glfw.GetTime()
}

func (w *Window) SinceLastFrame() float64 {
	return w.dt
}
