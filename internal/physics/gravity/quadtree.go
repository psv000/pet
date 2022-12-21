package gravity

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

const (
	NW NodeLocation = iota
	NE
	SW
	SE
	RootNode   NodeLocation = -1
	InvalidLoc NodeLocation = -2
)

var (
	Theta = 0.9
	Gamma float64
)

type (
	Particle interface {
		Mass() float64
		Position() mgl64.Vec3
	}

	TreeNode struct {
		quad [4]*TreeNode

		measuring Measuring
		loc       NodeLocation

		mass       float64
		center     mgl64.Vec3
		subdivided bool

		particle     Particle
		particlesNum int

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
		loc: RootNode,
	}
}

func (n *TreeNode) Place(objs []Particle) {
	for _, obj := range objs {
		n.Insert(obj)
	}
}

func (n *TreeNode) Locate(point mgl64.Vec3) NodeLocation {
	p0 := n.measuring.p0
	p1 := n.measuring.p1
	c := n.measuring.center

	if p0[0] < point[0] && point[0] <= c[0] {
		if c[1] < point[1] && point[1] <= p1[1] {
			return NW
		} else if p0[1] < point[1] && point[1] <= c[1] {
			return SW
		}
	}
	if c[0] < point[0] && point[0] <= p1[0] {
		if c[1] < point[1] && point[1] <= p1[1] {
			return NE
		} else if p0[1] < point[1] && point[1] <= c[1] {
			return SE
		}
	}
	return InvalidLoc
}

func (n *TreeNode) Quadrant(loc NodeLocation) {
	p0 := n.measuring.p0
	p1 := n.measuring.p1
	c := n.measuring.center

	switch loc {
	case NW:
		n.quad[NW] = &TreeNode{measuring: Measuring{p0: mgl64.Vec3{p0[0], c[1], 0.}, p1: mgl64.Vec3{c[0], p1[1], 0.}}, level: n.level + 1}
	case NE:
		n.quad[NE] = &TreeNode{measuring: Measuring{p0: mgl64.Vec3{c[0], c[1], 0.}, p1: mgl64.Vec3{p1[0], p1[1], 0.}}, level: n.level + 1}
	case SW:
		n.quad[SW] = &TreeNode{measuring: Measuring{p0: mgl64.Vec3{p0[0], p0[1], 0.}, p1: mgl64.Vec3{c[0], c[1], 0.}}, level: n.level + 1}
	case SE:
		n.quad[SE] = &TreeNode{measuring: Measuring{p0: mgl64.Vec3{c[0], p0[1], 0.}, p1: mgl64.Vec3{p1[0], c[1], 0.}}, level: n.level + 1}
	default:
		panic(true)
	}
	m := n.quad[loc].measuring
	n.quad[loc].measuring.center = m.p0.Add(m.p1).Mul(0.5)
}

func (n *TreeNode) Insert(p Particle) {
	if n.particlesNum > 1 {
		loc := n.Locate(p.Position())
		if loc != InvalidLoc {
			if n.quad[loc] == nil {
				n.Quadrant(loc)
			}
			n.quad[loc].Insert(p)
		}
	} else if n.particlesNum == 1 {
		{
			if n.particle == nil {
				panic("where is the particle?")
			}
			loc := n.Locate(n.particle.Position())
			if n.quad[loc] == nil {
				n.Quadrant(loc)
			}
			n.quad[loc].Insert(n.particle)
			n.particle = nil
		}
		{
			loc := n.Locate(p.Position())
			if n.quad[loc] == nil {
				n.Quadrant(loc)
			}
			n.quad[loc].Insert(p)
		}
	} else if n.particlesNum == 0 {
		n.particle = p
	}
	n.particlesNum++
}

func (n *TreeNode) MassDistribution() {
	if n.particlesNum == 1 {
		if n.particle == nil {
			panic("where is the particle?")
		}
		n.mass = n.particle.Mass()
		n.center = n.particle.Position()
		return
	}

	n.mass = 0.
	n.center = mgl64.Vec3{0., 0., 0.}

	for _, node := range n.quad {
		if node == nil {
			continue
		}
		node.MassDistribution()
		n.mass += node.mass
		n.center = n.center.Add(node.center.Mul(node.mass))
	}
	n.center = n.center.Mul(1. / n.mass)
}

func (n *TreeNode) GravityForce(p Particle) (val mgl64.Vec3) {
	val = n.gravityForceTree(p).Mul(p.Mass())
	return
}

func (n *TreeNode) gravityForceTree(p Particle) (val mgl64.Vec3) {
	var (
		r, k, d float64
	)

	if n.particlesNum == 1 {
		val = acceleration(p, n.particle)
		return
	}

	vec := p.Position().Sub(n.center)
	r = vec.Dot(vec)
	d = n.measuring.p1.Sub(n.measuring.p0).X()
	n.subdivided = d/r > Theta
	if n.subdivided {
		if n.quad[0] == nil {
			return
		}
		for _, node := range n.quad {
			if node == nil {
				continue
			}
			val = val.Add(node.gravityForceTree(p))
		}
	} else {
		k = Gamma * n.mass / (r * r * r)
		val = n.center.Sub(p.Position()).Mul(k)
	}
	return
}

func (n *TreeNode) Clear() {
	for i := range n.quad {
		n.quad[i] = nil
	}
	n.center = mgl64.Vec3{}
	n.mass = 0.
	n.particlesNum = 0
	n.particle = nil
}

func acceleration(p0, p1 Particle) (val mgl64.Vec3) {
	if p0 == p1 {
		return
	}

	r := p1.Position().Sub(p0.Position())
	rsq := r.Dot(r)

	if rsq > 0 {
		d := math.Sqrt(rsq)
		k := Gamma * p1.Mass() / (d * d * d)
		val = r.Mul(k)
	}
	return
}
