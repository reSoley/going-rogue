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

	view := NewView()

	return &Engine{
		view: view,
	}
}

func (e *Engine) Run() error {
	defer termbox.Close()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}

			e.view.processKeyInput(ev.Ch)
		}
	}

	return nil
}
