package main

import (
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type Application struct {
	name           string
	camera         *Camera
	context        *Context
	masterRenderer *RenderMaster
	states         []State
	isPopState     bool
}

func NewApplication(name string) *Application {
	app := &Application{
		name:           name,
		camera:         NewCamera(),
		context:        NewContext(),
		masterRenderer: NewRenderMaster(),
		states:         []State{},
	}
	app.pushState(NewStatePlaying(app))
	return app
}

func (a *Application) runLoop() {
	lastFrameTime := time.Now()
	for !a.context.window.ShouldClose() && len(a.states) > 0 {
		frameStart := time.Now()
		dt := frameStart.Sub(lastFrameTime)

		state := a.states[len(a.states)-1]
		state.handleInput()
		state.update(dt)
		a.camera.update()
		state.render(a.masterRenderer)

		a.masterRenderer.finishRender(a.context.window, a.camera)

		a.handleEvents()

		if a.isPopState {
			a.isPopState = false
			a.states = a.states[:len(a.states)-1]
			if len(a.states) > 0 {
				a.context.window.SetKeyCallback(a.states[len(a.states)-1].onKey)
			}
		}

		time.Sleep(15 * time.Millisecond)
	}
}

func (a *Application) getCamera() *Camera {
	return a.camera
}

func (a *Application) getWindow() *glfw.Window {
	return a.context.window
}

func (a *Application) handleEvents() {
	glfw.PollEvents()
}

func (a *Application) popState() {
	a.isPopState = true
}

func (a *Application) pushState(state State) {
	a.states = append(a.states, state)
	a.context.window.SetKeyCallback(state.onKey)
}

func (a *Application) dispose() {
	a.masterRenderer.dispose()
	a.context.dispose()
}
