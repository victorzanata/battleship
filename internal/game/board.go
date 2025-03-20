package game

import "fmt"

const BoardSize = 10

type Cell struct {
	HasShip bool
	Hit     bool
}

type Board struct {
	Grid           [BoardSize][BoardSize]Cell
	ShipsRemaining int
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) Display(hideShips bool) {
	fmt.Print("   ")
	for i := 0; i < BoardSize; i++ {
		fmt.Printf(" %d ", i)
	}
	fmt.Println()

	for row := 0; row < BoardSize; row++ {
		fmt.Printf("%2d ", row)
		for col := 0; col < BoardSize; col++ {
			cell := b.Grid[row][col]

			if cell.Hit {
				if cell.HasShip {
					fmt.Print("\x1b[31m ● \x1b[0m") // 🔴 Acertou um navio (Vermelho)
				} else {
					fmt.Print("\x1b[36m · \x1b[0m") // 🔵 Tiro na água (Azul)
				}
			} else if hideShips {
				fmt.Print("   ")
			} else if cell.HasShip {
				fmt.Print(" ■ ") // 🚢 Navio do jogador
			} else {
				fmt.Print(" □ ") // 🌊 Água visível no tabuleiro do jogador
			}
		}
		fmt.Println()
	}
}

func (b *Board) AlreadyShot(row, col int) bool {
	return b.Grid[row][col].Hit
}
