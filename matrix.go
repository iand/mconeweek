package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

func makeModelMatrix(entity Entity) mgl32.Mat4 {
	rot := entity.rotation()
	pos := entity.position()

	matrix := mgl32.HomogRotate3DX(mgl32.DegToRad(rot.X()))
	matrix = matrix.Mul4(mgl32.HomogRotate3DY(mgl32.DegToRad(rot.Y())))
	matrix = matrix.Mul4(mgl32.HomogRotate3DZ(mgl32.DegToRad(rot.Z())))
	matrix = matrix.Mul4(mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()))
	return matrix
}

func makeViewMatrix(camera *Camera) mgl32.Mat4 {
	rot := camera.rotation()
	pos := camera.position()

	matrix := mgl32.HomogRotate3DX(mgl32.DegToRad(rot.X()))
	matrix = matrix.Mul4(mgl32.HomogRotate3DY(mgl32.DegToRad(rot.Y())))
	matrix = matrix.Mul4(mgl32.HomogRotate3DZ(mgl32.DegToRad(rot.Z())))
	matrix = matrix.Mul4(mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()))
	return matrix
}

func makeProjectionMatrix(fov float32) mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(fov), 1280.0/720.0, 0.1, 1000.0)
}
