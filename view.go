package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type View struct {
	screenBuffer [viewWidth][viewHeight]rune
}

func newView() *View {
	var screenBuffer [viewWidth][viewHeight]rune

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
	for x, yBuffer := range v.screenBuffer {
		for y, _ := range yBuffer {
			termbox.SetCell(x, y, v.screenBuffer[x][y], termbox.ColorWhite, termbox.ColorDefault)
		}
	}

	err := termbox.Flush()
	if err != nil {
		return fmt.Errorf("Error flushing termbox buffer: %s", err)
	}

	return nil
}
