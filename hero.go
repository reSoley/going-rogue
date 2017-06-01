package main

import (
	"github.com/nsf/termbox-go"
)

type Hero struct {
	xPosition int
	yPosition int
}

func newHero(xPosition, yPosition int) *Hero {
	termbox.SetCell(xPosition, yPosition, '@', termbox.ColorWhite, termbox.ColorDefault)

	return &Hero{
		xPosition: xPosition,
		yPosition: yPosition,
	}
}

func (h *Hero) move(direction int) {
	termbox.SetCell(h.xPosition, h.yPosition, ' ', termbox.ColorWhite, termbox.ColorDefault)

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
}
