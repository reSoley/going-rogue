package main

import (
	"github.com/nsf/termbox-go"
)

type Hero struct {
	xPosition int
	yPosition int
}

func NewHero() *Hero {
	return &Hero{
		xPosition: 0,
		yPosition: 0,
	}
}

func (h *Hero) Move(direction int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	switch direction {
	case 0:
		h.xPosition--
	case 1:
		h.yPosition--
	case 2:
		h.xPosition++
	case 3:
		h.yPosition++
	}

	termbox.SetCell(h.xPosition, h.yPosition, '@', termbox.ColorWhite, termbox.ColorDefault)
	termbox.Flush()
}
