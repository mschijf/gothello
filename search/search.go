package search

import (
	"fmt"
	"gothello/board"
)

func ComputeMove(board board.HumanBoard) (col, row int) {
	GlobalString = ""
	computedMove := PnSearch(board.GetBitBoard(), board.GetColorToMove())
	col, row = computedMove.ToColRow()
	GlobalString += fmt.Sprintf("\n\nMove played: %c%c", 'A'+col, '1'+row)
	return col, row
}
