package main

import ()

type Creature struct {
	x          int
	y          int
	standingOn rune
	viewSlice  [][viewHeight]rune
}

func newCreature(x, y int, standingOn rune, viewSlice [][viewHeight]rune) *Creature {
	viewSlice[x][y] = '@'

	return &Creature{
		x:          x,
		y:          y,
		standingOn: standingOn,
		viewSlice:  viewSlice,
	}
}

func (c *Creature) moveLeft() {
	if c.x > 0 {
		c.viewSlice[c.x][c.y] = c.standingOn
		c.x--
		c.standingOn = c.viewSlice[c.x][c.y]
		c.viewSlice[c.x][c.y] = '@'
	}
}

func (c *Creature) moveUp() {
	if c.y > 0 {
		c.viewSlice[c.x][c.y] = c.standingOn
		c.y--
		c.standingOn = c.viewSlice[c.x][c.y]
		c.viewSlice[c.x][c.y] = '@'
	}
}

func (c *Creature) moveRight() {
	if c.x < viewWidth-1 {
		c.viewSlice[c.x][c.y] = c.standingOn
		c.x++
		c.standingOn = c.viewSlice[c.x][c.y]
		c.viewSlice[c.x][c.y] = '@'
	}
}

func (c *Creature) moveDown() {
	if c.y < viewHeight-1 {
		c.viewSlice[c.x][c.y] = c.standingOn
		c.y++
		c.standingOn = c.viewSlice[c.x][c.y]
		c.viewSlice[c.x][c.y] = '@'
	}
}
