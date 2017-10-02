package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Entity interface {
	position() mgl32.Vec3
	rotation() mgl32.Vec3
}

type SimpleEntity struct {
	pos, rot mgl32.Vec3
}

func (s *SimpleEntity) position() mgl32.Vec3 {
	return s.pos
}

func (s *SimpleEntity) rotation() mgl32.Vec3 {
	return s.rot
}
