package search

import (
	"gothello/board"
)

func ComputeMove(board board.HumanBoard) (col, row int) {
	computedMove := PnSearch(board.GetBitBoard(), board.GetColorToMove())
	return computedMove.ToColRow()
}
