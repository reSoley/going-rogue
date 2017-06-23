package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

const (
	ViewWidth  = 20
	ViewHeight = 10
)

type View struct {
	screenBuffer [ViewWidth][ViewHeight]rune
}

func newView() *View {
	var buffer [ViewWidth][ViewHeight]rune

	for x, yBuffer := range screenBuffer {
		for y, _ := range yBuffer {
			screenBuffer[x][y] = ' '
		}
	}

	return &View{
		screenBuffer: screenBuffer,
	}
}

func (v *View) render() error {
	for x, yBuffer := range v.buffer {
		for y, cur := range yBuffer {
			termbox.SetCell(x, y, cur, termbox.ColorWhite, termbox.ColorDefault)
		}
	}

	err := termbox.Flush()
	if err != nil {
		return fmt.Errorf("Error flushing termbox buffer: %s", err)
	}

	return nil
}
