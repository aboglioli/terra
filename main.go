package main

import (
	"fmt"
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	core.InitEngine("Terra", 1024, 768)
	defer core.DestroyEngine()

	events := core.NewEventPool()

	var x, y int32 = 0, 0

	events.AddKeyboardEvent(func(e *sdl.KeyboardEvent) {
		switch e.Keysym.Sym {
		case sdl.K_a:
			x -= 64
		case sdl.K_d:
			x += 64
		case sdl.K_w:
			y -= 64
		case sdl.K_s:
			y += 64
		}
	})

	renderer := core.Renderer()

	loop := core.NewLoop(60, func(d float64) {
		// Clear screen
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		// Draw
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.FillRect(&sdl.Rect{x, y, 64, 64})

		// Update
		renderer.Present()

		// Handle events
		events.HandleEvents()
	})

	events.AddKeyboardEvent(func(e *sdl.KeyboardEvent) {
		if e.Keysym.Sym == sdl.K_ESCAPE {
			loop.Stop()
		}
	})

	events.AddQuitEvent(func() {
		loop.Stop()
	})

	events.AddMouseMotionEvent(func(t *sdl.MouseMotionEvent) {
		fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
	})

	events.AddMouseButtonEvent(func(t *sdl.MouseButtonEvent) {
		fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
	})

	loop.Start()
}
