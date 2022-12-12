package solid

import (
	"github.com/go-gl/mathgl/mgl32"
	"pet/internal/app"
	"pet/internal/physics"
	"unsafe"

	"pet/internal/graphics"
	"pet/internal/window"
)

type Scene struct {
	window  *window.Window
	camera  *graphics.Camera
	objects map[uintptr]*Object

	ph *physics.System
}

func NewScene(window *window.Window, ph *physics.System) *Scene {
	return &Scene{
		window:  window,
		camera:  graphics.NewCamera(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, -1, 0}, -90, 0, window.InputManager()),
		objects: make(map[uintptr]*Object, 32),
		ph:      ph,
	}
}

func (s *Scene) ObtainObject() *Object {
	obj := NewObject(
		app.NewDot(),
		s.ph.ObtainObject(),
	)
	s.objects[uintptr(unsafe.Pointer(obj))] = obj
	return obj
}

func (s *Scene) ReleaseObject(obj *Object) {
	delete(s.objects, uintptr(unsafe.Pointer(obj)))
}

func (s *Scene) Update(dt float64, speed float64) {
	s.ph.Update(dt * speed)
	s.camera.Update(dt)

	fov := float32(60.0)
	projection := mgl32.Perspective(mgl32.DegToRad(fov),
		float32(s.window.Width())/float32(s.window.Height()),
		0.1,
		100.0)
	camTransform := s.camera.GetTransform()

	for i := range s.objects {
		obj := s.objects[i]
		obj.Update(dt, projection, camTransform)
	}
}
