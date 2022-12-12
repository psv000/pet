package main

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
	solid2 "pet/internal/app/solid"
	"pet/internal/resources"
	"runtime"
	"time"

	"pet/internal/physics"
	"pet/internal/window"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	var err error

	err = glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window := window.NewWindow(1024, 768, "Pet")
	resources.Load()

	physicsSystem := physics.NewSystem()
	scene := solid2.NewScene(window, physicsSystem)

	setup(window, scene)
}

func setup(win *window.Window, scene *solid2.Scene) {
	dot1 := scene.ObtainObject()
	dot1.M = 100_000.
	scl := float32(3.)
	dot1.SetScale(scl)
	dot1.SetColor(mgl32.Vec4{241. / 255., 120. / 255., 41. / 255., 1.})
	dot1.R = 0.1 * float64(scl)

	lenf := 38
	objs := make([]*solid2.Object, lenf)
	for i := 0; i < lenf; i++ {
		obj := scene.ObtainObject()
		obj.SetColor(mgl32.Vec4{0., 191. / 255., 1., 1.})
		obj.M = 300.
		scl := float32(obj.M / 1000)
		obj.SetScale(scl)
		obj.R = 0.1 * float64(scl)
		objs[i] = obj
	}

	reset := func() {
		dot1.SetPosition(mgl32.Vec3{})
		dot1.SetVelocity(mgl64.Vec3{})

		width := 6
		start := float32(-2.)
		var pos mgl32.Vec3
		pos[0] = start
		pos[1] = 0.8
		for i, obj := range objs {
			if i%width == 0 {
				pos[1] += 0.2
				pos[0] = start
			} else {
				pos[0] += 5. / float32(width)
			}
			obj.SetPosition(pos)
			obj.V = mgl64.Vec3{}
		}
	}

	reset()
	speedFactor := 1.
	const (
		speedStep = 0.01
		speedMin  = speedStep
		speedMax  = 3.
	)
	for !win.ShouldClose() {
		win.BeforeRender()
		if win.InputManager().IsActive(window.ProgramPause) {
			time.Sleep(time.Millisecond * 16)
			win.ResetFrameTime()
			continue
		}
		if win.InputManager().IsActive(window.ProgramReset) {
			reset()
		}

		if win.InputManager().IsActive(window.Faster) {
			speedFactor += speedStep
			speedFactor = math.Min(speedFactor, speedMax)
		}
		if win.InputManager().IsActive(window.Slower) {
			speedFactor -= speedStep
			speedFactor = math.Max(speedFactor, speedMin)
		}

		win.Render(func(dt float64) {
			//dot2.V = mgl64.Vec3{}
			scene.Update(dt, speedFactor)
		})
	}
}
