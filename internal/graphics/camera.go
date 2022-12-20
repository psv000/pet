package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"pet/internal/window"

	"github.com/go-gl/mathgl/mgl64"
)

type Camera struct {
	// Camera options
	moveSpeed         float64
	cursorSensitivity float64

	// Eular Angles
	pitch float64
	yaw   float64

	// Camera attributes
	pos     mgl32.Vec3
	front   mgl32.Vec3
	up      mgl32.Vec3
	right   mgl32.Vec3
	worldUp mgl32.Vec3

	inputManager *window.InputManager
}

func NewCamera(position, worldUp mgl32.Vec3, yaw, pitch float64, im *window.InputManager) *Camera {
	cam := Camera{
		moveSpeed:         5.00,
		cursorSensitivity: 0.1,
		pitch:             pitch,
		yaw:               yaw,
		pos:               position,
		up:                mgl32.Vec3{0, 1, 0},
		worldUp:           worldUp,
		inputManager:      im,
	}

	cam.updateDirection()
	return &cam
}

func (c *Camera) Update(dt float64) {
	c.updatePosition(dt)
	if c.inputManager.IsMouseButtonPushed() {
		c.updateDirection()
	}
}

func (c *Camera) updatePosition(dt float64) {
	adjustedSpeed := float32(dt * c.moveSpeed)

	if c.inputManager.IsActive(window.CameraUp) {
		c.pos = c.pos.Add(c.up.Mul(adjustedSpeed))
	}
	if c.inputManager.IsActive(window.CameraDown) {
		c.pos = c.pos.Sub(c.up.Mul(adjustedSpeed))
	}
	if c.inputManager.IsActive(window.CameraLeft) {
		c.pos = c.pos.Sub(c.front.Cross(c.up).Normalize().Mul(adjustedSpeed))
	}
	if c.inputManager.IsActive(window.CameraRight) {
		c.pos = c.pos.Add(c.front.Cross(c.up).Normalize().Mul(adjustedSpeed))
	}
	if c.inputManager.IsActive(window.CameraForward) {
		c.pos = c.pos.Add(c.front.Mul(adjustedSpeed))
	}
	if c.inputManager.IsActive(window.CameraBackward) {
		c.pos = c.pos.Sub(c.front.Mul(adjustedSpeed))
	}
	if c.inputManager.IsActive(window.ProgramReset) {
		//c.pos = mgl32.Vec3{0., 0., 3.}
		//c.pitch = 0
		//c.yaw = -90
		//
		//c.updateDirection()
	}
}

func (c *Camera) updateDirection() {
	dCursor := c.inputManager.CursorChange()

	dx := -c.cursorSensitivity * dCursor[0]
	dy := c.cursorSensitivity * dCursor[1]

	c.pitch += dy
	if c.pitch > 89.0 {
		c.pitch = 89.0
	} else if c.pitch < -89.0 {
		c.pitch = -89.0
	}

	c.yaw = math.Mod(c.yaw+dx, 360)
	c.updateVectors()
}

func (c *Camera) updateVectors() {
	// x, y, z
	c.front[0] = float32(math.Cos(mgl64.DegToRad(c.pitch)) * math.Cos(mgl64.DegToRad(c.yaw)))
	c.front[1] = float32(math.Sin(mgl64.DegToRad(c.pitch)))
	c.front[2] = float32(math.Cos(mgl64.DegToRad(c.pitch)) * math.Sin(mgl64.DegToRad(c.yaw)))
	c.front = c.front.Normalize()

	// Gram-Schmidt process to figure out right and up vectors
	c.right = c.worldUp.Cross(c.front).Normalize()
	c.up = c.right.Cross(c.front).Normalize()
}

func (c *Camera) GetTransform() mgl32.Mat4 {
	cameraTarget := c.pos.Add(c.front)

	return mgl32.LookAt(
		c.pos.X(), c.pos.Y(), c.pos.Z(),
		cameraTarget.X(), cameraTarget.Y(), cameraTarget.Z(),
		c.up.X(), c.up.Y(), c.up.Z(),
	)
}
