package main

import (
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type StatePlaying struct {
	player    *Player
	app       *Application
	keyStates keyStates
}

func NewStatePlaying(app *Application) *StatePlaying {
	s := &StatePlaying{
		player: NewPlayer(mgl32.Vec3{0, 0, -5}),
		app:    app,
	}
	app.getCamera().hookEntity(s.player)
	return s
}

func (s *StatePlaying) handleEvent() {}

func (s *StatePlaying) handleInput() {
	s.player.handleInput(s.app.context.window, s.keyStates)
}

func (s *StatePlaying) update(dt time.Duration) {
	s.player.update(dt)
}

func (s *StatePlaying) render(renderer *RenderMaster) {
	renderer.drawQuad(mgl32.Vec3{0, 0, 0})
}

func (s *StatePlaying) onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		switch key {
		case glfw.KeyEscape:
			s.app.context.window.SetShouldClose(true)
		case glfw.KeyW:
			s.keyStates.w = true
		case glfw.KeyA:
			s.keyStates.a = true
		case glfw.KeyS:
			s.keyStates.s = true
		case glfw.KeyD:
			s.keyStates.d = true
		}
	case glfw.Release:
		switch key {
		case glfw.KeyW:
			s.keyStates.w = false
		case glfw.KeyA:
			s.keyStates.a = false
		case glfw.KeyS:
			s.keyStates.s = false
		case glfw.KeyD:
			s.keyStates.d = false
		}
	}
}

type keyStates struct {
	w, a, s, d bool
}
