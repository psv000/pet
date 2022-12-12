package collisions

import "github.com/go-gl/mathgl/mgl64"

type PlaneCollided interface {
	Collided
	Position() mgl64.Vec3
	Size() mgl64.Vec3
}

func SphereVsPlane(sphere SphereCollided, plane PlaneCollided) bool {
	return false
}
