package main

type GoingRogue struct {
	viewSlice [][viewHeight]rune
}

func newGoingRogue(viewSlice [][viewHeight]rune) *GoingRogue {
	return &GoingRogue{
		viewSlice: viewSlice,
	}
}

func (g *GoingRogue) update() {
	for x, yBuf := range g.viewSlice {
		for y, _ := range yBuf {
			g.viewSlice[x][y] = '.'
		}
	}
}
