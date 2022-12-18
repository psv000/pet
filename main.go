package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
	solid "pet/internal/app/solid"
	"pet/internal/physics"
	"pet/internal/resources"
	"pet/internal/window"
	"runtime"
	"time"
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
	for _, p := range resources.ProgramList {
		p.OnResolutionChange(mgl32.Vec2{1024, 768})
	}

	physicsSystem := physics.NewSystem()
	scene := solid.NewScene(window, physicsSystem)

	setup(window, scene)
}

func setup(win *window.Window, scene *solid.Scene) {
	dot1 := scene.ObtainSphere()
	dot1.Physics().SetMass(2.)
	scl := float32(3.)
	dot1.Graphics().SetScale(scl)
	dot1.Graphics().SetColor(mgl32.Vec4{241. / 255., 120. / 255., 41. / 255., 1.})
	dot1.Physics().SetRadius(0.1 * float64(scl))
	dot1.Physics().SetRestitution(0.9)

	dot2 := scene.ObtainSphere()
	dot2.Physics().SetMass(3.)
	scl = float32(2.7)
	dot2.Graphics().SetScale(scl)
	dot2.Graphics().SetColor(mgl32.Vec4{137. / 255., 18. / 255., 89. / 255., 1.})
	dot2.Physics().SetRadius(0.1 * float64(scl))
	dot2.Physics().SetRestitution(0.5)

	var objs []*solid.Sphere
	for i := 0; i < 10; i++ {
		obj := scene.ObtainSphere()
		obj.Physics().SetMass(0.5 + 0.05*float64(i))
		scl = float32(0.4)
		obj.Graphics().SetScale(scl)
		obj.Graphics().SetColor(mgl32.Vec4{182. / 255., 181. / 255., 233. / 255., 1.})
		obj.Physics().SetRadius(0.1 * float64(scl))
		obj.Physics().SetRestitution(0.6)
		objs = append(objs, obj)
	}

	reset := func() {
		dot1.SetPosition(mgl32.Vec3{})
		dot1.Physics().SetVelocity(mgl64.Vec3{0.5, 0, 0})

		dot2.SetPosition(mgl32.Vec3{1.1, -0.3, 0})
		dot2.Physics().SetVelocity(mgl64.Vec3{-0.5, -0.1, 0})

		for i, obj := range objs {
			obj.Physics().SetPosition(mgl64.Vec3{-0.8 + 0.2*float64(i), 1, 0})
			obj.Physics().SetVelocity(mgl64.Vec3{})
		}
	}

	pos := []mgl64.Vec3{
		{-2, -1.5, 0},
		{-2, -1.4, 0},
		{-2, 1.5, 0},
		{1.9, -1.4, 0},
	}
	size := []mgl64.Vec3{
		{4, 0.1, 0},
		{0.1, 3, 0},
		{4, 0.1, 0},
		{0.1, 3, 0},
	}

	for i := 0; i < len(pos); i++ {
		r := scene.ObtainWall(pos[i], size[i])
		r.Graphics().SetColor(mgl32.Vec4{0.3, 0.01, 0.1, 1})
		r.Physics().SetMass(9999)
		r.Physics().SetRestitution(0.4)
		r.Physics().SetSize(size[i])
		r.Physics().SetPosition(pos[i])
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
			dot1.Physics().SetVelocity(mgl64.Vec3{0.})

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
			scene.Update(dt, speedFactor)
		})
	}
}
