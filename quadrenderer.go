package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type QuadRenderer struct {
	quads []mgl32.Vec3

	quadModel    *Model
	shader       *BasicShader
	basicTexture *BasicTexture
}

func NewQuadRenderer() *QuadRenderer {
	q := &QuadRenderer{
		shader:       NewBasicShader("BasicVertex.glsl", "BasicFragment.glsl"),
		basicTexture: NewBasicTexture("Test"),
		quadModel: NewModel(

			[]float32{
				-0.5, 0.5, 0,
				0.5, 0.5, 0,
				0.5, -0.5, 0,
				-0.5, -0.5, 0,
			},
			[]float32{
				0, 1,
				1, 1,
				1, 0,
				0, 0,
			},
			[]uint32{
				0, 1, 2,
				2, 3, 0,
			}),
	}
	return q
}

func (q *QuadRenderer) add(pos mgl32.Vec3) {
	q.quads = append(q.quads, pos)
}

func (q *QuadRenderer) renderQuads(camera *Camera) {
	q.shader.useProgram()
	q.quadModel.bindVAO()
	q.basicTexture.bindTexture()
	q.shader.loadProjectionViewMatrix(camera.getProjectionViewMatrix())

	for _, quad := range q.quads {
		mm := makeModelMatrix(&SimpleEntity{pos: quad})
		q.shader.loadModelMatrix(mm)
		gl.DrawElements(gl.TRIANGLES, q.quadModel.getIndicesCount(), gl.UNSIGNED_INT, nil)
	}

	q.quads = q.quads[:0]
}

func (q *QuadRenderer) dispose() {
	q.basicTexture.dispose()
	q.shader.dispose()
	q.quadModel.dispose()
}
