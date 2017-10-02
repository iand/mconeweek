package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	id uint32
}

func NewShader(vertexFile string, fragmentFile string) *Shader {
	s := Shader{
		id: loadShaders(vertexFile, fragmentFile),
	}
	return &s
}

func (s *Shader) useProgram() {
	gl.UseProgram(s.id)
}

func (s *Shader) dispose() {
	gl.DeleteProgram(s.id)
}

func (s *Shader) loadInt(location int32, value int32) {
	gl.Uniform1i(location, value)
}

func (s *Shader) loadFloat(location int32, value float32) {
	gl.Uniform1f(location, value)
}

func (s *Shader) loadVector2(location int32, vect mgl32.Vec2) {
	gl.Uniform2f(location, vect.X(), vect.Y())
}

func (s *Shader) loadVector3(location int32, vect mgl32.Vec3) {
	gl.Uniform3f(location, vect.X(), vect.Y(), vect.Z())
}

func (s *Shader) loadVector4(location int32, vect mgl32.Vec4) {
	gl.Uniform4f(location, vect.X(), vect.Y(), vect.Z(), vect.W())
}

func (s *Shader) loadMatrix4(location int32, matrix mgl32.Mat4) {
	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
}
