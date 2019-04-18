package main

import (
	"fmt"
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	core.InitSystems("Terra", 1024, 768)
	defer core.DestroySystems()

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

	// state := &core.GameState{
	// 	Video:  video,
	// 	Events: events,
	// }

	surface := core.Surface()

	loop := core.NewLoop(60, func(d float64) {
		surface.FillRect(nil, 0)

		rect := sdl.Rect{x, y, 64, 64}
		surface.FillRect(&rect, 0xffff0000)
		surface.FillRect(&sdl.Rect{x, y, 32, 32}, 0xff00ff00)

		core.Update()
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
