package main

type GoingRogue struct {
	hero *Creature
}

func newGoingRogue(viewSlice [][viewHeight]rune, controller *Controller) *GoingRogue {
	floor := newFloor(viewSlice)
	floor.drawRooms()
	hero := newCreature(viewWidth/2, viewHeight/2, '.', viewSlice)

	controller.bind('a', hero.moveLeft)
	controller.bind('w', hero.moveUp)
	controller.bind('d', hero.moveRight)
	controller.bind('s', hero.moveDown)

	return &GoingRogue{
		hero: hero,
	}
}
