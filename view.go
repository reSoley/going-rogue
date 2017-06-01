package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type View struct {
	hero       *Hero
	termWidth  int
	termHeight int
}

func newView() *View {
	width, height := termbox.Size()
	hero := newHero(width/2, height/2)

	return &View{
		hero:       hero,
		termWidth:  width,
		termHeight: height,
	}
}

func (v *View) render() error {
	err := termbox.Flush()
	if err != nil {
		return fmt.Errorf("Error flushing termbox buffer: %s", err)
	}

	return nil
}

func (v *View) processKeyInput(input rune) error {
	switch input {
	case 'a':
		v.hero.move(0)
	case 'w':
		v.hero.move(1)
	case 'd':
		v.hero.move(2)
	case 's':
		v.hero.move(3)
	}

	err := v.render()
	if err != nil {
		return fmt.Errorf("Error processing keystroke input: %s", err)
	}

	return nil
}
