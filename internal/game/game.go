package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Play() {
	playerBoard := NewBoard()
	opponentBoard := NewBoard()

	shipSizes := []int{5, 4, 3, 3, 2}
	for _, size := range shipSizes {
		playerBoard.PlaceShipRandomly(size)
		opponentBoard.PlaceShipRandomly(size)
	}

	for {
		fmt.Println("\n\x1b[1;37mYour Board:\x1b[0m")
		playerBoard.Display(false)
		fmt.Println("\n\x1b[1;37mOpponent's Board:\x1b[0m")
		opponentBoard.Display(false)

		row, col := GetPlayerInput(opponentBoard)
		result := opponentBoard.Fire(row, col)
		fmt.Println("You fired at", row, col, "and it was a", result)

		if IsGameOver(opponentBoard) {
			fmt.Println("\x1b[1;32mCongratulations! You won!\x1b[0m")
			break
		}

		ox, oy := GenerateOpponentMove(playerBoard)
		result = playerBoard.Fire(ox, oy)
		fmt.Println("Opponent fired at", ox, oy, "and it was a", result)

		if IsGameOver(playerBoard) {
			fmt.Println("\x1b[1;31mOh no! Your ships were destroyed!\x1b[0m")
			break
		}
	}
}

func GetPlayerInput(b *Board) (int, int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\x1b[1;37mEnter coordinates to fire (row, col): \x1b[0m")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		coords := strings.Split(input, ",")
		if len(coords) != 2 {
			fmt.Println("\x1b[1;31mInvalid input. Use format: row,col\x1b[0m")
			continue
		}

		row, err1 := strconv.Atoi(strings.TrimSpace(coords[0]))
		col, err2 := strconv.Atoi(strings.TrimSpace(coords[1]))

		if err1 != nil || err2 != nil || row < 0 || row >= BoardSize || col < 0 || col >= BoardSize {
			fmt.Println("\x1b[1;31mInvalid coordinates. Try again.\x1b[0m")
			continue
		}

		if b.AlreadyShot(row, col) {
			fmt.Println("\x1b[1;33mYou already fired at this position! Choose another.\x1b[0m")
			continue
		}

		return row, col
	}
}

func GenerateOpponentMove(b *Board) (int, int) {
	for {
		x, y := rand.Intn(BoardSize), rand.Intn(BoardSize)
		if !b.AlreadyShot(x, y) {
			return x, y
		}
	}
}

func IsGameOver(b *Board) bool {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if b.Grid[i][j].HasShip && !b.Grid[i][j].Hit {
				return false
			}
		}
	}
	return true
}
