package collisions

import (
	"pet/internal/physics/model"

	"github.com/go-gl/mathgl/mgl64"
)

const (
	NW NodeLocation = iota
	NE
	SW
	SE
	InvalidLoc NodeLocation = -2
	RootNode   NodeLocation = -1
)

const (
	objectsInQuad = 2
	deep          = 4
)

type (
	TreeNode struct {
		quad [4]*TreeNode

		measuring Measuring
		loc       NodeLocation

		objects []model.Model

		level int
	}

	NodeLocation int

	Measuring struct {
		center mgl64.Vec3
		p0, p1 mgl64.Vec3
	}
)

func NewTreeNode(p0, p1 mgl64.Vec3) *TreeNode {
	return &TreeNode{
		measuring: Measuring{
			p0: p0, p1: p1,
			center: p0.Add(p1).Mul(0.5),
		},
		loc:   RootNode,
		level: 1,
	}
}

func (n *TreeNode) Retrieve(obj model.Model) (out []model.Model) {
	loc := n.Locate(obj)
	if n.quad[0] != nil && loc != InvalidLoc {
		out = n.quad[loc].Retrieve(obj)
	}

	out = append(out, n.objects...)
	return
}

func (n *TreeNode) Split() {
	c := n.measuring.p1.Add(n.measuring.p0).Mul(0.5)
	p0 := n.measuring.p0
	p1 := n.measuring.p1

	n.quad[NW] = &TreeNode{measuring: Measuring{p0: mgl64.Vec3{p0[0], c[1], 0.}, p1: mgl64.Vec3{c[0], p1[1], 0.}}}
	n.quad[NE] = &TreeNode{measuring: Measuring{p0: mgl64.Vec3{c[0], c[1], 0.}, p1: mgl64.Vec3{p1[0], p1[1], 0.}}}
	n.quad[SW] = &TreeNode{measuring: Measuring{p0: mgl64.Vec3{p0[0], p0[1], 0.}, p1: mgl64.Vec3{c[0], c[1], 0.}}}
	n.quad[SE] = &TreeNode{measuring: Measuring{p0: mgl64.Vec3{c[0], p0[1], 0.}, p1: mgl64.Vec3{p1[0], c[1], 0.}}}

	for i, node := range n.quad {
		m := node.measuring
		node.measuring.center = m.p0.Add(m.p1).Mul(0.5)
		node.loc = NodeLocation(i)
		node.level = n.level + 1
	}
}

func (n *TreeNode) Locate(obj model.Model) NodeLocation {
	rect := obj.Rect()

	c := n.measuring.center
	p0 := n.measuring.p0
	p1 := n.measuring.p1

	if p0[0] < rect[0] && rect[2] <= c[0] {
		if c[1] < rect[1] && rect[3] <= p1[1] {
			return NW
		} else if p0[1] < rect[1] && rect[3] <= c[1] {
			return SW
		}
	}
	if c[0] < rect[0] && rect[2] <= p1[0] {
		if c[1] < rect[1] && rect[3] <= p1[1] {
			return NE
		} else if p0[1] < rect[1] && rect[3] <= c[1] {
			return SE
		}
	}
	return InvalidLoc
}

func (n *TreeNode) Insert(obj model.Model) {
	if n.quad[0] != nil {
		loc := n.Locate(obj)
		if loc != InvalidLoc {
			n.quad[loc].Insert(obj)
			return
		}
	}
	n.objects = append(n.objects, obj)

	if len(n.objects) > objectsInQuad && n.level <= deep {
		if n.quad[0] == nil {
			n.Split()
		}

		move := make([]model.Model, len(n.objects))
		copy(move, n.objects)
		var i int

		for _, mv := range move {
			loc := n.Locate(mv)
			if loc == InvalidLoc {
				i++
				continue
			}
			n.quad[loc].Insert(mv)
			n.objects = append(n.objects[0:i], n.objects[i+1:len(n.objects)]...)
		}
	}
}

func (n *TreeNode) Clear() {
	if n.quad[0] != nil {
		for _, n := range n.quad {
			n.Clear()
		}
	}
	for i := range n.quad {
		n.quad[i] = nil
	}
	n.level = -1
	n.objects = nil
}
