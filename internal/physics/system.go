package physics

import (
	"github.com/go-gl/mathgl/mgl64"
	"pet/internal/physics/collisions"
	"sort"
)

type System struct {
	g       float64
	objects []*Object
}

func NewSystem() *System {
	return &System{
		objects: make([]*Object, 0, 32),
		g:       6.6743e-7,
	}
}

func (s *System) ObtainObject() *Object {
	obj := &Object{}
	s.objects = append(s.objects, obj)
	return obj
}

func (s *System) ReleaseObject(del *Object) {
	for i, obj := range s.objects {
		if obj == del {
			s.objects[i] = s.objects[len(s.objects)-1]
			s.objects = s.objects[:len(s.objects)-1]
			return
		}
	}
}

func (s *System) Update(dt float64) {
	s.gravity(dt)
	s.collisions(dt)

	for _, obj1 := range s.objects {
		obj1.Update(dt)
	}
}

func (s *System) gravity(dt float64) {
	set := make(map[*Object]*Object, len(s.objects))
	for _, obj1 := range s.objects {
		for _, obj2 := range s.objects {
			if obj1 == obj2 || set[obj2] == obj1 {
				continue
			}
			set[obj1] = obj2
			Dir := obj2.P.Sub(obj1.P)
			S := Dir.Len()
			if S < 1e-3 {
				continue
			}
			F := s.g * (obj1.M * obj2.M) / (S * S)

			ƒ := Dir.Normalize().Mul(F)

			obj1.ApplyForce(dt, ƒ)
			ƒ1 := ƒ.Mul(-1.)
			obj2.ApplyForce(dt, ƒ1)
		}
	}
}

func (s *System) collisions(dt float64) {
	set := make(map[*Object]*Object, len(s.objects))

	var c mgl64.Vec3
	for _, obj := range s.objects {
		c = c.Add(obj.P)
	}
	c = c.Mul(1. / float64(len(s.objects)))

	sort.Slice(s.objects, func(i, j int) bool {
		p1 := s.objects[i].P.Sub(c)
		p2 := s.objects[i].P.Sub(c)
		if p1.Dot(p1)-p2.Dot(p2) < 1e-3 {
			if p1.X()-p2.X() < 1e3 {
				return p1.Y() >= p2.Y()
			}
			return p1.X() >= p2.X()
		}
		return p1.Dot(p1) < p2.Dot(p2)
	})

	for _, obj1 := range s.objects {
		for _, obj2 := range s.objects {
			if obj1 == obj2 || set[obj2] == obj1 {
				continue
			}
			set[obj1] = obj2

			if ok, correction := correction(obj1, obj2); ok {
				obj1.P = obj1.P.Add(correction)
				obj2.P = obj2.P.Sub(correction)

				ƒ1, ƒ2 := collisions.ResolveSpheres(obj1, obj2, dt)
				obj1.ApplyForce(dt, ƒ1)
				obj2.ApplyForce(dt, ƒ2)
			}
			for _, obj3 := range s.objects {
				if obj2 == obj3 || obj2 == obj3 {
					continue
				}

				if ok, correction := correction(obj3, obj2); ok {
					obj3.P = obj3.P.Add(correction)
					obj2.P = obj2.P.Sub(correction)

					ƒ1, ƒ2 := collisions.ResolveSpheres(obj3, obj2, dt)
					obj3.ApplyForce(dt, ƒ1)
					obj2.ApplyForce(dt, ƒ2)
				}
			}
		}
	}
}

func correction(obj1, obj2 *Object) (bool, mgl64.Vec3) {
	var collision bool
	var penetration float64
	if collision, penetration = collisions.Spheres(obj1, obj2); !collision {
		return false, mgl64.Vec3{}
	}
	normal := obj1.Position().Sub(obj2.Position()).Normalize()
	correction := normal.Mul(penetration / (obj1.M + obj2.M))
	return true, correction
}
