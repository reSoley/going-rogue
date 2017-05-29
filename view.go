package main

import (
	"github.com/nsf/termbox-go"
)

type View struct {
	hero       *Hero
	termWidth  int
	termHeight int
}

func NewView() *View {
	width, height := termbox.Size()
	hero := NewHero(width/2, height/2)

	return &View{
		hero:       hero,
		termWidth:  width,
		termHeight: height,
	}
}

func (v *View) Render() {
	termbox.Flush()
}

func (v *View) processKeyInput(input rune) {
	switch input {
	case 'a':
		v.hero.Move(0)
	case 'w':
		v.hero.Move(1)
	case 'd':
		v.hero.Move(2)
	case 's':
		v.hero.Move(3)
	}

	v.Render()
}
