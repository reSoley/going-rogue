package main

type Tile struct {
	cur  rune
	last rune
}

func newTile() *Tile {
	return &Tile{
		cur:  '.',
		last: '.',
	}
}

func (t *Tile) set(next rune) {
	t.last = t.cur
	t.cur = next
}

func (t *Tile) revert() {
	t.cur = t.last
}
