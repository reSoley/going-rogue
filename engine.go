package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

type Engine struct {
	view *View
}

func NewEngine() *Engine {
	err := termbox.Init()
	if err != nil {
		fmt.Printf("Failed to initialize termbox-go: %s", err)
		os.Exit(1)
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	view := newView()

	return &Engine{
		view: view,
	}
}

func (e *Engine) Run() error {
	defer termbox.Close()
	e.view.render()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}

			err := e.view.processKeyInput(ev.Ch)
			if err != nil {
				return fmt.Errorf("Error running engine: %s", err)
			}
		}
	}

	return nil
}
