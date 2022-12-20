package quadtree

import (
	"github.com/go-gl/mathgl/mgl64"
	"pet/internal/physics/model"
)

const (
	objectsInQuad = 2
	deep          = 4
)

type (
	Tree struct {
		root Quad
	}
)

func NewTree(lb, rt mgl64.Vec3) *Tree {
	return &Tree{
		root: Quad{
			bounds: bounds{
				lb: lb, rt: rt,
			},
			level: 1,
		},
	}
}

func (t *Tree) Insert(obj model.Model) {
	t.root.Insert(obj)
}

func (t *Tree) Retrieve(obj model.Model) (out []model.Model) {
	return t.root.Retrieve(obj)
}

func (t *Tree) Clear() {
	t.root.Clear()
}

func (t *Tree) Quads() []*Quad {
	return t.root.Quads()
}

func (q *Quad) Quads() []*Quad {
	if q.leaves[0] == nil {
		return []*Quad{q}
	}
	result := make([]*Quad, 0)
	if len(q.objects) > 0 {
		result = append(result, q)
	}
	for _, l := range q.leaves {
		result = append(result, l.Quads()...)
	}
	return result
}

func (q *Quad) Level() int {
	return q.level
}
