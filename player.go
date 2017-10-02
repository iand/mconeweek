package main

import (
	"math"
	"sync"
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Player struct {
	pos, rot, velocity mgl32.Vec3
	lastMousePosition  mgl32.Vec2
	mouseInit          sync.Once
}

func NewPlayer(pos mgl32.Vec3) *Player {
	return &Player{
		pos: pos,
	}
}

func (p *Player) handleInput(window *glfw.Window, keys keyStates) {
	p.keyboardInput(keys)
	p.mouseInput(window)
}

func (p *Player) update(dt time.Duration) {
	p.pos = p.pos.Add(p.velocity.Mul(float32(dt.Seconds())))
	p.velocity = p.velocity.Mul(0.95)
}

func (p *Player) mouseInput(window *glfw.Window) {
	const bound = 80
	p.mouseInit.Do(func() {
		x, y := window.GetCursorPos()
		p.lastMousePosition[0] = float32(x)
		p.lastMousePosition[1] = float32(y)
	})

	x, y := window.GetCursorPos()
	change := mgl32.Vec2{
		float32(x) - p.lastMousePosition[0],
		float32(y) - p.lastMousePosition[1],
	}

	p.rot[1] += change[0] * 0.05
	p.rot[0] += change[1] * 0.05

	if p.rot[0] > bound {
		p.rot[0] = bound
	} else if p.rot[0] < -bound {
		p.rot[0] = -bound
	}

	if p.rot[1] > 360 {
		p.rot[1] = 0
	} else if p.rot[1] < 0 {
		p.rot[1] = 360
	}

	cx, cy := window.GetSize()
	window.SetCursorPos(float64(cx)/4, float64(cy)/4)

	x, y = window.GetCursorPos()
	p.lastMousePosition[0] = float32(x)
	p.lastMousePosition[1] = float32(y)

}

func (p *Player) keyboardInput(keys keyStates) {
	var change mgl32.Vec3
	const speed = 0.01

	if keys.w {
		change[0] = float32(math.Cos(float64(mgl32.DegToRad(p.rot.Y()+90))) * speed)
		change[2] = float32(math.Sin(float64(mgl32.DegToRad(p.rot.Y()+90))) * speed)
	}

	if keys.s {
		change[0] = -float32(math.Cos(float64(mgl32.DegToRad(p.rot.Y()+90))) * speed)
		change[2] = -float32(math.Sin(float64(mgl32.DegToRad(p.rot.Y()+90))) * speed)
	}

	if keys.a {
		change[0] = float32(math.Cos(float64(mgl32.DegToRad(p.rot.Y()))) * speed)
		change[2] = float32(math.Sin(float64(mgl32.DegToRad(p.rot.Y()))) * speed)
	}

	if keys.d {
		change[0] = -float32(math.Cos(float64(mgl32.DegToRad(p.rot.Y()))) * speed)
		change[2] = -float32(math.Sin(float64(mgl32.DegToRad(p.rot.Y()))) * speed)
	}

	p.velocity = p.velocity.Add(change)

}

func (p *Player) position() mgl32.Vec3 {
	return p.pos
}

func (p *Player) rotation() mgl32.Vec3 {
	return p.rot
}
