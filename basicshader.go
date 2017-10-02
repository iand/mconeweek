package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type BasicShader struct {
	*Shader
	locationProjectionViewMatrix int32
	locationModelMatrix          int32
}

func NewBasicShader(vertexFile string, fragmentFile string) *BasicShader {
	b := &BasicShader{
		Shader: NewShader(vertexFile, fragmentFile),
	}
	b.getUniforms()
	return b
}

func (b *BasicShader) loadProjectionViewMatrix(pvMatrix mgl32.Mat4) {
	b.loadMatrix4(b.locationProjectionViewMatrix, pvMatrix)
}

func (b *BasicShader) loadModelMatrix(matrix mgl32.Mat4) {
	b.loadMatrix4(b.locationModelMatrix, matrix)
}

func (b *BasicShader) getUniforms() {
	b.useProgram()
	b.locationProjectionViewMatrix = gl.GetUniformLocation(b.id, gl.Str("projViewMatrix\x00"))
	b.locationModelMatrix = gl.GetUniformLocation(b.id, gl.Str("modelMatrix\x00"))
}
