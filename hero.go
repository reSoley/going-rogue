package main

import ()

type Hero struct {
	xPosition int
	yPosition int
}

func newHero(x, y int) *Hero {
	return &Hero{
		xPosition: x,
		yPosition: y,
	}
}

func (h *Hero) move(direction rune) {
	switch direction {
	case 'a':
		if h.xPosition > 0 {
			h.xPosition--
		}
	case 'w':
		if h.yPosition > 0 {
			h.yPosition--
		}
	case 'd':
		if h.xPosition < ViewWidth-1 {
			h.xPosition++
		}
	case 's':
		if h.yPosition < ViewHeight-1 {
			h.yPosition++
		}
	}
}
