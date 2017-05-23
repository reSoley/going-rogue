package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type Engine struct {
	hero Hero
}

func NewEngine() *Engine {
	hero := NewHero()

	return &Engine{
		hero: *hero,
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

	termbox.SetCell(e.hero.xPosition, e.hero.yPosition, '@', termbox.ColorWhite, termbox.ColorDefault)
	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
			if ev.Key == termbox.KeyArrowLeft {
				e.hero.Move(0)
			}
			if ev.Key == termbox.KeyArrowUp {
				e.hero.Move(1)
			}
			if ev.Key == termbox.KeyArrowRight {
				e.hero.Move(2)
			}
			if ev.Key == termbox.KeyArrowDown {
				e.hero.Move(3)
			}

			termbox.Flush()
		}
	}

	return nil
}
