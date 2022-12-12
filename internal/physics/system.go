package physics

import "pet/internal/physics/collisions"

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
	for _, obj1 := range s.objects {
		for _, obj2 := range s.objects {
			if obj1 == obj2 || set[obj2] == obj1 {
				continue
			}
			set[obj1] = obj2

			var collision bool
			var penetration float64
			if collision, penetration = collisions.Spheres(obj1, obj2); !collision {
				continue
			}

			normal := obj1.Position().Sub(obj2.Position())
			percent := 0.0008 / (obj1.M + obj2.M)

			correction := normal.Mul(penetration / (1./obj1.M + 1./obj2.M) * percent)
			obj1.P = obj1.P.Add(correction.Mul(obj1.M))
			obj2.P = obj2.P.Sub(correction.Mul(obj2.M))
			ƒ1, ƒ2 := collisions.ResolveSpheres(obj1, obj2, dt)

			obj1.ApplyForce(dt, ƒ1)
			obj2.ApplyForce(dt, ƒ2)
		}
	}
}
