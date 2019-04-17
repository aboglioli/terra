package main

import (
	// "fmt"
	"github.com/aboglioli/terra/core"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	video := core.InitVideo("Terra", 800, 600)
	defer video.Destroy()
	events := core.NewEventPool()

	var x, y int32 = 0, 0
	events.AddKeyboardEvent(sdl.K_a, func() {
		x -= 64
	})
	events.AddKeyboardEvent(sdl.K_d, func() {
		x += 64
	})
	events.AddKeyboardEvent(sdl.K_w, func() {
		y -= 64
	})
	events.AddKeyboardEvent(sdl.K_s, func() {
		y += 64
	})

	// state := &core.GameState{
	// 	Video:  video,
	// 	Events: events,
	// }

	loop := core.NewLoop(60, func(d float64) {
		video.Surface.FillRect(nil, 0)
		rect := sdl.Rect{x, y, 64, 64}
		video.Surface.FillRect(&rect, 0xffff0000)

		video.Update()

		events.HandleEvents()
		// fmt.Println(d)
	})

	events.AddKeyboardEvent(sdl.K_ESCAPE, func() {
		go loop.Stop()
	})

	events.AddQuitEvent(func() {
		go loop.Stop()
	})

	loop.Start()
}
