package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

const (
	viewWidth  = 20
	viewHeight = 10
)

type Engine struct {
	controller *Controller
	view       *View
	goingRogue *GoingRogue
}

func newEngine() *Engine {
	err := termbox.Init()
	if err != nil {
		fmt.Printf("Failed to initialize termbox-go: %s", err)
		os.Exit(1)
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	controller := newController()
	view := newView()
	goingRogue := newGoingRogue(view.screenBuffer[:], controller)

	return &Engine{
		controller: controller,
		view:       view,
		goingRogue: goingRogue,
	}
}

func (e *Engine) run() error {
	defer termbox.Close()
	e.view.render()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}

			e.controller.processKey(ev.Ch)
			e.view.render()
		}
	}

	return nil
}

func main() {
	e := newEngine()
	e.run()
}
