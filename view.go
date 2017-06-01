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
	buffer     [ViewWidth][ViewHeight]*Tile
	hero       *Hero
	termWidth  int
	termHeight int
}

func newView() *View {
	var buffer [ViewWidth][ViewHeight]*Tile
	for x, yBuffer := range buffer {
		for y, _ := range yBuffer {
			buffer[x][y] = newTile()
		}
	}

	hero := newHero(ViewWidth/2, ViewHeight/2)
	buffer[hero.xPosition][hero.yPosition].set('@')

	width, height := termbox.Size()

	return &View{
		buffer:     buffer,
		hero:       hero,
		termWidth:  width,
		termHeight: height,
	}
}

func (v *View) render() error {
	for x, yBuffer := range v.buffer {
		for y, tile := range yBuffer {
			termbox.SetCell(x, y, tile.cur, termbox.ColorWhite, termbox.ColorDefault)
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
		v.buffer[v.hero.xPosition][v.hero.yPosition].revert()
		v.hero.move(input)
		v.buffer[v.hero.xPosition][v.hero.yPosition].set('@')
	}

	err := v.render()
	if err != nil {
		return fmt.Errorf("Error processing keystroke input: %s", err)
	}

	return nil
}
