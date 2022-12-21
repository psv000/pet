package collisions

import (
	"pet/internal/physics/model"

	"github.com/go-gl/mathgl/mgl64"
)

func (n *TreeNode) Value() (mgl64.Vec3, mgl64.Vec3) {
	return n.measuring.p0, n.measuring.p1
}

func (n *TreeNode) Objects() []model.Model {
	return n.objects
}

func (n *TreeNode) Nodes() []*TreeNode {
	if n.quad[0] == nil {
		return []*TreeNode{n}
	}
	result := make([]*TreeNode, 0)
	if len(n.objects) > 0 {
		result = append(result, n)
	}
	for _, node := range n.quad {
		result = append(result, node.Nodes()...)
	}
	return result
}

func (n *TreeNode) Level() int {
	return n.level
}
