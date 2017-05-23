package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type Engine struct {
	cursorX int
	cursorY int
}

func NewEngine() *Engine {
	return &Engine{
		cursorX: 0,
		cursorY: 0,
	}
}

func (e *Engine) Run() error {
	err := termbox.Init()
	if err != nil {
		return fmt.Errorf("Failed to initialize termbox-go: %s", err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	termbox.SetCell(e.cursorX, e.cursorY, '@', termbox.ColorWhite, termbox.ColorDefault)
	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
			if ev.Key == termbox.KeyArrowUp {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				e.cursorY--
				termbox.SetCell(e.cursorX, e.cursorY, '@', termbox.ColorWhite, termbox.ColorDefault)
			}
			if ev.Key == termbox.KeyArrowDown {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				e.cursorY++
				termbox.SetCell(e.cursorX, e.cursorY, '@', termbox.ColorWhite, termbox.ColorDefault)
			}
			if ev.Key == termbox.KeyArrowLeft {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				e.cursorX--
				termbox.SetCell(e.cursorX, e.cursorY, '@', termbox.ColorWhite, termbox.ColorDefault)
			}
			if ev.Key == termbox.KeyArrowRight {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				e.cursorX++
				termbox.SetCell(e.cursorX, e.cursorY, '@', termbox.ColorWhite, termbox.ColorDefault)
			}

			termbox.Flush()
		}
	}

	return nil
}
