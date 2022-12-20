package quadtree

import (
	"github.com/go-gl/mathgl/mgl64"
	"pet/internal/physics/model"
)

const (
	QuadLoc1 = iota
	QuadLoc2
	QuadLoc3
	QuadLoc4
	QuadLocInvalid = -1
)

func alloc(quad Quad) *Quad {
	ptr := &Quad{}
	*ptr = quad
	return ptr
}

type (
	Quad struct {
		leaves [4]*Quad

		bounds  bounds
		objects []model.Model

		level int
	}

	QuadLoc int

	bounds struct {
		lb, rt mgl64.Vec3
	}
)

func (q *Quad) Retrieve(obj model.Model) (out []model.Model) {
	loc := q.Find(obj)
	if q.leaves[0] != nil && loc != QuadLocInvalid {
		out = q.leaves[loc].Retrieve(obj)
	}

	out = append(out, q.objects...)
	return
}

func (q *Quad) Split() {
	c := q.bounds.rt.Add(q.bounds.lb).Mul(0.5)
	lb := q.bounds.lb
	rt := q.bounds.rt

	q.leaves[QuadLoc1] = alloc(Quad{bounds: bounds{lb: mgl64.Vec3{lb[0], c[1], 0.}, rt: mgl64.Vec3{c[0], rt[1], 0.}}, level: q.level + 1})
	q.leaves[QuadLoc2] = alloc(Quad{bounds: bounds{lb: mgl64.Vec3{c[0], c[1], 0.}, rt: mgl64.Vec3{rt[0], rt[1], 0.}}, level: q.level + 1})
	q.leaves[QuadLoc3] = alloc(Quad{bounds: bounds{lb: mgl64.Vec3{lb[0], lb[1], 0.}, rt: mgl64.Vec3{c[0], c[1], 0.}}, level: q.level + 1})
	q.leaves[QuadLoc4] = alloc(Quad{bounds: bounds{lb: mgl64.Vec3{c[0], lb[1], 0.}, rt: mgl64.Vec3{rt[0], c[1], 0.}}, level: q.level + 1})
}

func (q *Quad) Find(obj model.Model) QuadLoc {
	rect := obj.Rect()
	c := q.bounds.rt.Add(q.bounds.lb).Mul(0.5)
	if rect[0] > q.bounds.lb[0] && rect[2] <= c[0] {
		if rect[1] > q.bounds.lb[1] && rect[3] <= c[1] {
			return QuadLoc3
		} else if rect[1] > c[1] && rect[3] <= q.bounds.rt[1] {
			return QuadLoc1
		}
	}
	if rect[0] > c[0] && rect[2] <= q.bounds.rt[0] {
		if rect[1] > q.bounds.lb[1] && rect[3] <= c[1] {
			return QuadLoc4
		} else if rect[1] > c[1] && rect[3] <= q.bounds.rt[1] {
			return QuadLoc2
		}
	}
	return QuadLocInvalid
}

func (q *Quad) Insert(obj model.Model) {
	if q.leaves[0] != nil {
		loc := q.Find(obj)
		if loc != QuadLocInvalid {
			q.leaves[loc].Insert(obj)
			return
		}
	}
	q.objects = append(q.objects, obj)

	if len(q.objects) > objectsInQuad && q.level <= deep {
		if q.leaves[0] == nil {
			q.Split()
		}

		move := make([]model.Model, len(q.objects))
		copy(move, q.objects)
		var i int

		for _, mv := range move {
			loc := q.Find(mv)
			if loc == QuadLocInvalid {
				i++
				continue
			}
			q.leaves[loc].Insert(mv)
			q.objects = append(q.objects[0:i], q.objects[i+1:len(q.objects)]...)
		}
	}
}

func (q *Quad) Clear() {
	if q.leaves[0] != nil {
		for _, l := range q.leaves {
			l.Clear()
		}
	}
	for i := range q.leaves {
		q.leaves[i] = nil
	}
	q.level = -1
	q.objects = nil
}

func (q *Quad) Value() (mgl64.Vec3, mgl64.Vec3) {
	return q.bounds.lb, q.bounds.rt
}

func (q *Quad) Objects() []model.Model {
	return q.objects
}
