package main

import (
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type State interface {
	handleEvent()
	handleInput()
	update(dt time.Duration)
	render(renderer *RenderMaster)
	onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
}
