package physics

import (
	"github.com/go-gl/mathgl/mgl64"
	"pet/internal/physics/collisions"
	"pet/internal/physics/gravity"
	"pet/internal/pkg/sampler"
)

type (
	System struct {
		g, a       float64
		resolution int

		trees

		spheres []collisions.Sphere
		cuboids []collisions.CuboidCollided
	}

	trees struct {
		collisions *collisions.TreeNode
		gravity    *gravity.TreeNode
	}
)

func NewSystem() *System {
	return &System{
		spheres: make([]collisions.Sphere, 0, 32),
		cuboids: make([]collisions.CuboidCollided, 0, 32),
		trees: trees{
			collisions: collisions.NewTreeNode(mgl64.Vec3{-2, -2, 0}, mgl64.Vec3{2, 2, 0}),
			gravity:    gravity.NewTreeNode(mgl64.Vec3{-2, -2, 0}, mgl64.Vec3{2, 2, 0}),
		},
		resolution: 20,
		g:          6.6743e-11,
		a:          0.981e1,
	}
}

func (s *System) Tree() *collisions.TreeNode {
	return s.trees.collisions
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

var (
	insertSampler   = sampler.New(4800, "insert tree")
	retrieveSampler = sampler.New(2400, "retrieve")
)

func (s *System) Update(dt float64) {
	dt /= float64(s.resolution)
	for i := 0; i < s.resolution; i++ {

		s.trees.collisions.Clear()
		s.trees.gravity.Clear()

		gravity.Gamma = s.g

		for _, obj := range s.spheres {
			s.trees.collisions.Insert(obj)
		}
		for _, obj := range s.spheres {
			s.trees.gravity.Insert(obj)
		}

		for _, obj := range s.cuboids {
			s.trees.collisions.Insert(obj)
		}

		s.gravity(dt)
		s.collisions(dt)

		for _, obj := range s.spheres {
			obj.Update(dt)
		}

	}
}

func (s *System) gravity(dt float64) {
	s.trees.gravity.MassDistribution()
	for _, obj := range s.spheres {
		ƒ := s.trees.gravity.GravityForce(obj)
		obj.ApplyForce(dt, ƒ)
		gravity.Global(s.a, dt, obj)
	}
}

func (s *System) collisions(dt float64) {
	for _, obj1 := range s.spheres {
		retrieves := s.trees.collisions.Retrieve(obj1)
		for _, obj2 := range retrieves {
			if obj1 == obj2 {
				continue
			}

			switch obj := obj2.(type) {
			case collisions.Sphere:
				var collision bool
				var penetration float64
				if collision, penetration = collisions.DetectSpheresCollision(obj1, obj); !collision {
					continue
				}

				collisions.AmendSpheres(penetration, obj1, obj)
				collisions.ResolveSpheresCollision(dt, obj1, obj)
			case collisions.CuboidCollided:
				ok, penetration, point := collisions.DetectSphereVsCuboidCollision(obj1, obj)
				if !ok {
					continue
				}

				sdt := obj1.Position().Sub(point).Normalize().Mul(penetration)
				obj1.SetPosition(obj1.Position().Add(sdt))

				collisions.ResolveSphereVsCuboidCollision(obj1, obj, point)
			}

		}
	}
}

func (s *System) Bounds() (lb, rt mgl64.Vec3) {
	lb = mgl64.Vec3{mgl64.MaxValue, mgl64.MaxValue, 0.}
	rt = mgl64.Vec3{mgl64.MinValue, mgl64.MinValue, 0.}
	for _, obj := range s.spheres {
		lb = Min(lb, obj.Position())
		rt = Max(rt, obj.Position())
	}
	return
}

func Min(l, r mgl64.Vec3) (out mgl64.Vec3) {
	out = r
	if l[0] < r[0] {
		out[0] = l[0]
	}
	if l[1] < r[1] {
		out[1] = l[1]
	}
	return
}

func Max(l, r mgl64.Vec3) (out mgl64.Vec3) {
	out = r
	if l[0] >= r[0] {
		out[0] = l[0]
	}
	if l[1] >= r[1] {
		out[1] = l[1]
	}
	return
}
