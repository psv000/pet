package solid

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"pet/internal/physics"
	"time"
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
		NewDot(),
		s.ph.ObtainObject(),
	)
	s.objects[uintptr(unsafe.Pointer(obj))] = obj
	return obj
}

func (s *Scene) ReleaseObject(obj *Object) {
	delete(s.objects, uintptr(unsafe.Pointer(obj)))
}

const (
	samplePerFrames = 240
)

var (
	phSamplerDur time.Duration
	phCounter    int

	gSamplerDur time.Duration
	gCounter    int
)

func sample(fu func(), frames int, counter *int, duration *time.Duration) {
	t := time.Now()
	fu()
	*duration += time.Since(t)
	*counter++
	c := *counter
	if c%frames == 0 {
		d := *duration
		fmt.Printf("%v\n", d/time.Duration(c))
		*counter = 0
		*duration = 0
	}
}

func (s *Scene) Update(dt float64, speed float64) {
	sample(func() {
		s.ph.Update(dt * speed)
	}, samplePerFrames, &phCounter, &phSamplerDur)
	sample(func() {
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
	}, samplePerFrames, &gCounter, &gSamplerDur)
}
