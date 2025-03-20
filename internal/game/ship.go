package game

import (
	"math/rand"
)

type Ship struct {
	Positions []Position
}

type Position struct {
	X, Y int
}

func (b *Board) PlaceShipRandomly(size int) {
	for {
		x, y := rand.Intn(BoardSize), rand.Intn(BoardSize)
		vertical := rand.Intn(2) == 0

		if b.CanPlaceShip(x, y, size, vertical) {
			for i := 0; i < size; i++ {
				if vertical {
					b.Grid[x][y+i].HasShip = true
				} else {
					b.Grid[x+i][y].HasShip = true
				}
			}
			b.ShipsRemaining++
			break
		}
	}
}

func (b *Board) CanPlaceShip(x, y, size int, vertical bool) bool {
	if vertical {
		if y+size > BoardSize {
			return false
		}
		for i := 0; i < size; i++ {
			if b.Grid[x][y+i].HasShip {
				return false
			}
		}
	} else {
		if x+size > BoardSize {
			return false
		}
		for i := 0; i < size; i++ {
			if b.Grid[x+i][y].HasShip {
				return false
			}
		}
	}
	return true
}

func (b *Board) ShipDestroyed(x, y int) bool {
	return true
}

func (b *Board) Fire(x, y int) string {
	if b.AlreadyShot(x, y) {
		return "Already Shot"
	}
	b.Grid[x][y].Hit = true
	if b.Grid[x][y].HasShip {
		if b.ShipDestroyed(x, y) {
			b.ShipsRemaining--
			return "Destroyed"
		}
		return "Hit"
	}
	return "Miss"
}
