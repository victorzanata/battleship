package game

// Define the size of the game board as a constant.
const BoardSize = 10

type CellState struct {
	Ship  bool
	Hit   bool
	Empty bool
	Miss  bool
}

type Board struct {
	Grid [BoardSize][BoardSize]CellState
}

func NewBoard() *Board {
	return &Board{}
}
