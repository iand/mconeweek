package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type RenderMaster struct {
	quadRenderer *QuadRenderer
}

func NewRenderMaster() *RenderMaster {
	return &RenderMaster{
		quadRenderer: NewQuadRenderer(),
	}
}

func (r *RenderMaster) drawQuad(pos mgl32.Vec3) {
	r.quadRenderer.add(pos)
}

func (r *RenderMaster) finishRender(window *glfw.Window, camera *Camera) {
	gl.ClearColor(0.1, 0.5, 1.0, 1.0)

	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	r.quadRenderer.renderQuads(camera)
	window.SwapBuffers()
}

func (r *RenderMaster) dispose() {
	r.quadRenderer.dispose()
}
