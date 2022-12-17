package physics

import (
	"github.com/go-gl/mathgl/mgl64"
)

type System struct {
	g, a float64

	spheres []SphereCollided
	cuboids []CuboidCollided
}

func NewSystem() *System {
	return &System{
		spheres: make([]SphereCollided, 0, 32),
		cuboids: make([]CuboidCollided, 0, 32),
		g:       6.6743e-3,
		a:       9.81,
	}
}

func (s *System) ObtainSphere() SphereCollided {
	sphere := NewSphere()
	s.spheres = append(s.spheres, sphere)
	return sphere
}

func (s *System) ReleaseSphere(del SphereCollided) {
	for i, obj := range s.spheres {
		if obj == del {
			s.spheres[i] = s.spheres[len(s.spheres)-1]
			s.spheres = s.spheres[:len(s.spheres)-1]
			return
		}
	}
}

func (s *System) ObtainCuboid() CuboidCollided {
	cuboid := NewCuboid()
	s.cuboids = append(s.cuboids, cuboid)
	return cuboid
}

func (s *System) ReleaseCuboid(del CuboidCollided) {
	for i, obj := range s.cuboids {
		if obj == del {
			s.cuboids[i] = s.cuboids[len(s.cuboids)-1]
			s.cuboids = s.cuboids[:len(s.cuboids)-1]
			return
		}
	}
}

func (s *System) Update(dt float64) {
	s.gravity(dt)
	s.collisions(dt)

	for _, obj := range s.spheres {
		obj.Update(dt)
	}
}

func (s *System) gravity(dt float64) {
	set := make(map[SphereCollided]SphereCollided, len(s.spheres))
	for _, obj1 := range s.spheres {
		for _, obj2 := range s.spheres {
			if obj1 == obj2 || set[obj2] == obj1 {
				continue
			}
			set[obj1] = obj2
			Dir := obj2.Position().Sub(obj1.Position())
			r := Dir.Len()
			r2 := r * r
			F := s.g * obj1.Mass() * obj2.Mass() / r2

			ƒ1 := Dir.Normalize().Mul(F)
			ƒ2 := ƒ1.Mul(-1.)

			obj1.ApplyForce(dt, ƒ1)
			obj2.ApplyForce(dt, ƒ2)
		}
		obj1.ApplyForce(dt, mgl64.Vec3{0., -1., 0.}.Mul(obj1.Mass()*s.a))
	}
}

func (s *System) collisions(dt float64) {
	processed := make(map[SphereCollided]SphereCollided, len(s.spheres))

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
			if collision, penetration = DetectSpheresCollision(obj1, obj2); !collision {
				continue
			}
			AmendSpheres(penetration, obj1, obj2)
			ResolveSpheresCollision(dt, obj1, obj2)
		}
	}
	for _, obj1 := range s.spheres {
		for _, obj2 := range s.cuboids {
			ok, penetration, point := DetectSphereVsCuboidCollision(obj1, obj2)
			if !ok {
				continue
			}
			sdt := obj1.Position().Sub(point).Normalize().Mul(penetration * dt)
			obj1.SetPosition(obj1.Position().Add(sdt))

			ResolveSphereVsCuboidCollision(obj1, obj2, point)
		}
	}
}
