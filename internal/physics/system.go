package physics

import (
	"github.com/go-gl/mathgl/mgl64"
	"pet/internal/physics/collisions"
	"pet/internal/physics/gravity"
)

type System struct {
	g, a       float64
	resolution int

	spheres []collisions.Sphere
	cuboids []collisions.CuboidCollided
}

func NewSystem() *System {
	return &System{
		spheres:    make([]collisions.Sphere, 0, 32),
		cuboids:    make([]collisions.CuboidCollided, 0, 32),
		resolution: 10,
		g:          6.6743e-3,
		a:          0.981e1,
	}
}

func (s *System) ObtainSphere() collisions.Sphere {
	sphere := collisions.NewSphere()
	s.spheres = append(s.spheres, sphere)
	return sphere
}

func (s *System) ReleaseSphere(del collisions.Sphere) {
	for i, obj := range s.spheres {
		if obj == del {
			s.spheres[i] = s.spheres[len(s.spheres)-1]
			s.spheres = s.spheres[:len(s.spheres)-1]
			return
		}
	}
}

func (s *System) ObtainCuboid() collisions.CuboidCollided {
	cuboid := collisions.NewCuboid()
	s.cuboids = append(s.cuboids, cuboid)
	return cuboid
}

func (s *System) ReleaseCuboid(del collisions.CuboidCollided) {
	for i, obj := range s.cuboids {
		if obj == del {
			s.cuboids[i] = s.cuboids[len(s.cuboids)-1]
			s.cuboids = s.cuboids[:len(s.cuboids)-1]
			return
		}
	}
}

func (s *System) Update(dt float64) {
	dt /= float64(s.resolution)
	for i := 0; i < s.resolution; i++ {
		s.gravity(dt)
		s.collisions(dt)

		for _, obj := range s.spheres {
			obj.Update(dt)
		}
	}
}

func (s *System) gravity(dt float64) {
	set := make(map[collisions.Sphere]collisions.Sphere, len(s.spheres))
	for _, obj1 := range s.spheres {
		for _, obj2 := range s.spheres {
			if obj1 == obj2 || set[obj2] == obj1 {
				continue
			}
			set[obj1] = obj2
			gravity.Apply(s.g, dt, obj1, obj2)
		}
		gravity.Global(s.a, dt, obj1)
	}
}

func (s *System) collisions(dt float64) {
	processed := make(map[collisions.Sphere]collisions.Sphere, len(s.spheres))

	var c mgl64.Vec3
	for _, obj := range s.spheres {
		c = c.Add(obj.Position())
	}
	c = c.Mul(1. / float64(len(s.spheres)))

	for _, obj1 := range s.spheres {
		for _, obj2 := range s.spheres {
			if obj1 == obj2 || processed[obj2] == obj1 {
				continue
			}
			processed[obj1] = obj2

			var collision bool
			var penetration float64
			if collision, penetration = collisions.DetectSpheresCollision(obj1, obj2); !collision {
				continue
			}
			collisions.AmendSpheres(penetration, obj1, obj2)
			collisions.ResolveSpheresCollision(dt, obj1, obj2)
		}
	}
	for _, obj1 := range s.spheres {
		for _, obj2 := range s.cuboids {
			ok, penetration, point := collisions.DetectSphereVsCuboidCollision(obj1, obj2)
			if !ok {
				continue
			}
			_ = penetration
			sdt := obj1.Position().Sub(point).Normalize().Mul(penetration)
			obj1.SetPosition(obj1.Position().Add(sdt))

			collisions.ResolveSphereVsCuboidCollision(obj1, obj2, point)
		}
	}
}
