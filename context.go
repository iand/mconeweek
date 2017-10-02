package main

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Context struct {
	window *glfw.Window
}

func NewContext() *Context {
	if err := glfw.Init(); err != nil {
		log.Fatalf("failed to initialize glfw: %v", err)
	}

	glfw.WindowHint(glfw.DepthBits, 24)
	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.Decorated, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(1280, 720, "Minecraft", nil, nil)
	if err != nil {
		log.Fatalf("failed to create window: %v", err)
	}
	window.MakeContextCurrent()
	window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	// Enable vsync (swap buffers every 1 frame), set to 0 for no vsync
	glfw.SwapInterval(0)

	if err := gl.Init(); err != nil {
		log.Fatalf("failed to initialize glow: %v", err)
	}
	gl.Viewport(0, 0, 1280, 720)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	return &Context{
		window: window,
	}

}

func (c *Context) dispose() {
	glfw.Terminate()
}
