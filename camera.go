package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	pos, rot         mgl32.Vec3
	entity           Entity
	projectionMatrix mgl32.Mat4
	viewMatrix       mgl32.Mat4
	projViewMatrix   mgl32.Mat4
}

func NewCamera() *Camera {
	return &Camera{
		projectionMatrix: makeProjectionMatrix(90),
		pos:              mgl32.Vec3{0, 0, -3.5},
	}
}

func (c *Camera) update() {
	if c.entity != nil {
		c.pos = c.entity.position()
		c.rot = c.entity.rotation()
	}
	c.viewMatrix = makeViewMatrix(c)
	c.projViewMatrix = c.projectionMatrix.Mul4(c.viewMatrix)
}

func (c *Camera) hookEntity(entity Entity) {
	c.entity = entity
}

func (c *Camera) getViewMatrix() mgl32.Mat4 {
	return c.viewMatrix
}

func (c *Camera) getProjMatrix() mgl32.Mat4 {
	return c.projectionMatrix
}

func (c *Camera) getProjectionViewMatrix() mgl32.Mat4 {
	return c.projViewMatrix
}

func (c *Camera) position() mgl32.Vec3 {
	return c.pos
}

func (c *Camera) rotation() mgl32.Vec3 {
	return c.rot
}
