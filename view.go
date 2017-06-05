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
	buffer     [ViewWidth][ViewHeight]rune
	floor      *Floor
	hero       *Hero
	termWidth  int
	termHeight int
}

func newView() *View {
	var buffer [ViewWidth][ViewHeight]rune
	for x, yBuffer := range buffer {
		for y, _ := range yBuffer {
			buffer[x][y] = '.'
		}
	}

	// floor := newFloor(buffer)
	// floor.drawRoom(5, 2, 6, 6)

	hero := newHero(ViewWidth/2, ViewHeight/2)
	buffer[hero.xPosition][hero.yPosition] = '@'

	width, height := termbox.Size()

	return &View{
		buffer: buffer,
		// floor:      floor,
		hero:       hero,
		termWidth:  width,
		termHeight: height,
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

func (v *View) processKeyInput(input rune) error {
	switch input {
	case 'a', 'w', 'd', 's':
		v.buffer[v.hero.xPosition][v.hero.yPosition] = v.hero.standingOn
		v.hero.move(input)
		v.hero.setStandingOn(v.buffer[v.hero.xPosition][v.hero.yPosition])
		v.buffer[v.hero.xPosition][v.hero.yPosition] = '@'
	}

	err := v.render()
	if err != nil {
		return fmt.Errorf("Error processing keystroke input: %s", err)
	}

	return nil
}
