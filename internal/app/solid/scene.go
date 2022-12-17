package solid

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"pet/internal/physics"
	"time"
	"unsafe"

	"pet/internal/graphics"
	"pet/internal/window"
)

type Scene struct {
	window     *window.Window
	camera     *graphics.Camera
	spheres    map[uintptr]*Sphere
	rectangles map[uintptr]*Wall

	ph *physics.System
}

func NewScene(window *window.Window, ph *physics.System) *Scene {
	return &Scene{
		window:     window,
		camera:     graphics.NewCamera(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, -1, 0}, -90, 0, window.InputManager()),
		spheres:    make(map[uintptr]*Sphere, 32),
		rectangles: make(map[uintptr]*Wall, 32),
		ph:         ph,
	}
}

func (s *Scene) ObtainSphere() *Sphere {
	obj := NewSphere(
		NewDot(),
		s.ph.ObtainSphere(),
	)
	s.spheres[uintptr(unsafe.Pointer(obj))] = obj
	return obj
}

func (s *Scene) ReleaseSphere(obj *Sphere) {
	delete(s.spheres, uintptr(unsafe.Pointer(obj)))
}

func (s *Scene) ObtainWall(position, size mgl64.Vec3) *Wall {
	obj := NewWall(
		NewRectangle(Vec3L(position), Vec3L(size)),
		s.ph.ObtainCuboid(),
	)
	s.rectangles[uintptr(unsafe.Pointer(obj))] = obj
	return obj
}

func (s *Scene) ReleaseWall(obj *Wall) {
	delete(s.rectangles, uintptr(unsafe.Pointer(obj)))
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

		for i := range s.spheres {
			obj := s.spheres[i]
			obj.Update(dt, projection, camTransform)
		}
		for i := range s.rectangles {
			obj := s.rectangles[i]
			obj.Update(dt, projection, camTransform)
		}
	}, samplePerFrames, &gCounter, &gSamplerDur)
}
