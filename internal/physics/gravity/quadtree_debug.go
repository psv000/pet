package gravity

import "github.com/go-gl/mathgl/mgl64"

func (n *TreeNode) Nodes() []*TreeNode {
	var out []*TreeNode
	if n.particlesNum > 0 {
		out = append(out, n)
	}
	for _, node := range n.quad {
		if node != nil {
			out = append(out, node.Nodes()...)
		}
	}
	return out
}

func (n *TreeNode) Rect() (mgl64.Vec3, mgl64.Vec3) {
	return n.measuring.p0, n.measuring.p1
}

func (n *TreeNode) Mass() float64 {
	return n.mass
}
