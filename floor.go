package main

import (
	"math/rand"
	"time"
)

const (
	floorSize      = 9
	maxNumRooms    = 7
	maxPathPercent = 9
	minPathPercent = 4
)

type Room struct {
	neighbors []int
}

func newRoom() *Room {
	var neighbors []int

	return &Room{
		neighbors: neighbors,
	}
}

type Floor struct {
	buffer [][viewHeight]rune
	rooms  [floorSize]*Room
}

func newFloor(buffer [][viewHeight]rune) *Floor {
	var rooms [floorSize]*Room
	var queue []int

	rand.Seed(time.Now().UTC().UnixNano())

	viewed := make(map[int]bool)
	start := rand.Intn(floorSize)
	queue = append(queue, start)

	pathPercent := maxPathPercent

	for numRooms := 0; len(queue) > 0 && numRooms < maxNumRooms; numRooms++ {
		cur := queue[0]
		queue = queue[1:]

		rooms[cur] = newRoom()
		viewed[cur] = true

		for _, validNeighbor := range getValidNeighbors(cur) {
			if _, has := viewed[validNeighbor]; !has && rand.Intn(10) < pathPercent {
				queue = append(queue, validNeighbor)
				rooms[cur].neighbors = append(rooms[cur].neighbors, validNeighbor)
			}

			if pathPercent > minPathPercent {
				pathPercent--
			}
		}
	}

	return &Floor{
		buffer: buffer,
		rooms:  rooms,
	}
}

func (f *Floor) drawRooms() {
	widthStep := viewWidth / 3
	heightStep := viewHeight / 3

	for i, room := range f.rooms {
		if room != nil {
			xPosition := (i / 3) * widthStep
			yPosition := (i % 3) * heightStep
			f.drawRoom(xPosition, yPosition, widthStep, heightStep)
		}
	}
}

// Not super readable, needs cleanup
// Possibly weight rand calls to make rooms favor larger sizes
func (f *Floor) drawRoom(xPosition, yPosition, maxWidth, maxHeight int) {
	roomWidth := randRange(4, maxWidth-1)
	roomHeight := randRange(4, maxHeight-1)
	roomXPosition := randRange(xPosition+1, xPosition+maxWidth-roomWidth)
	roomYPosition := randRange(yPosition+1, yPosition+maxHeight-roomHeight)

	for i := roomXPosition; i < roomXPosition+roomWidth; i++ {
		for j := roomYPosition; j < roomYPosition+roomHeight; j++ {
			if j == roomYPosition || j == roomYPosition+roomHeight-1 {
				f.buffer[i][j] = '-'
			} else if i == roomXPosition || i == roomXPosition+roomWidth-1 {
				f.buffer[i][j] = '|'
			} else {
				f.buffer[i][j] = '.'
			}
		}
	}
}

// May be more useful in some uitlities file
func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func getValidNeighbors(room int) []int {
	var validNeighbors []int

	switch room {
	case 0:
		validNeighbors = []int{1, 3}
	case 1:
		validNeighbors = []int{0, 2, 4}
	case 2:
		validNeighbors = []int{1, 5}
	case 3:
		validNeighbors = []int{0, 4, 6}
	case 4:
		validNeighbors = []int{1, 3, 5, 7}
	case 5:
		validNeighbors = []int{2, 4, 8}
	case 6:
		validNeighbors = []int{3, 7}
	case 7:
		validNeighbors = []int{4, 6, 8}
	case 8:
		validNeighbors = []int{5, 7}
	}

	return validNeighbors
}
