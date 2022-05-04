package board

import (
	"errors"
	"fmt"
	"gothello/bit64math"
	"strconv"
	"strings"
)

//           BitBoard                               Human (0-based) Board
//
//  --- --- --- --- --- --- --- ---    RIJ --- --- --- --- --- --- --- ---
// |63 |62 |61 |60 |59 |58 |57 |56 |     0|   |   |   |   |   |   |   |   |
//  --- --- --- --- --- --- --- ---        --- --- --- --- --- --- --- ---
// |55 |   |   |   |   |   |   |48 |     1|   |   |   |   |   |   |   |   |
//  --- --- --- --- --- --- --- ---        --- --- --- --- --- --- --- ---
// |47 |   |   |   |   |   |   |40 |     2|   |   |   |   |   |   |   |   |
//  --- --- --- --- --- --- --- ---        --- --- --- --- --- --- --- ---
// |39 |   |   |   |   |   |   |32 |     3|   |   |   |   |   |   |   |   |
//  --- --- --- --- --- --- --- ---        --- --- --- --- --- --- --- ---
// |31 |   |   |   |   |   |   |24 |     4|   |   |   |   |   |   |   |   |
//  --- --- --- --- --- --- --- ---        --- --- --- --- --- --- --- ---
// |23 |   |   |   |   |   |   |16 |     5|   |   |   |   |   |   |   |   |
//  --- --- --- --- --- --- --- ---        --- --- --- --- --- --- --- ---
// |15 |   |   |   |   |   |   | 8 |     6|   |   |   |   |   |   |   |   |
//  --- --- --- --- --- --- --- ---        --- --- --- --- --- --- --- ---
// | 7 | 6 | 5 | 4 | 3 | 2 | 1 | 0 |     7|   |   |   |   |   |   |   |   |
//  --- --- --- --- --- --- --- ---        --- --- --- --- --- --- --- ---
//                                          0   1   2   3   4   5   6   7   KOLOM

const DELIMITER = ":"

func colRowToBit(col, row int) uint64 {
	return 1 << ((7-row)*8 + (7 - col))
}

func bitToColRow(bit uint64) (int, int) {
	mrb := bit64math.MostRightBitIndex(bit)
	return 7 - (mrb % 8), 7 - (mrb / 8)
}

func (bb *BitBoard) IsBlackToMove() bool {
	return bb.ColorToMove == BLACK
}

func (bb *BitBoard) IsWhiteDisc(col, row int) bool {
	return colRowToBit(col, row)&bb.Board[WHITE] != 0
}

func (bb *BitBoard) IsBlackDisc(col, row int) bool {
	return colRowToBit(col, row)&bb.Board[BLACK] != 0
}

func (bb *BitBoard) IsPlayable(col, row int) bool {
	return colRowToBit(col, row)&bb.GetAllCandidateMoves() != 0
}

func (bb *BitBoard) MustPass() bool {
	return bb.GetAllCandidateMoves() == 0
}

func (bb *BitBoard) HasHistory() bool {
	return !bb.Stack.isEmpty()
}

func (bb *BitBoard) doBitBoardMove(moveBit uint64) error {
	var moves = bb.GenerateMoves()
	for _, move := range moves {
		if move.discPlayed == moveBit {
			bb.DoMove(&move)
			return nil
		}
	}
	return errors.New("move from UI is not correct")
}

func (bb *BitBoard) DoColRowMove(col, row int) {
	bb.doBitBoardMove(colRowToBit(col, row))
}

func (bb *BitBoard) DoPassMove() {
	bb.doBitBoardMove(0)
}

func (bb *BitBoard) ToBoardString() string {
	return fmt.Sprintf("%d%s%d%s%d", bb.Board[0], DELIMITER, bb.Board[1], DELIMITER, bb.ColorToMove)
}

func (bb *BitBoard) ToBoardStatusString() string {
	var movesPlayedString = ""
	var tmpStack MoveStack
	for !bb.Stack.isEmpty() {
		move := bb.Stack.fromTop(0)
		tmpStack.push(move)
		if move.isPass() {
			movesPlayedString = "88" + movesPlayedString
		} else {
			col, row := bitToColRow(move.discPlayed)
			movesPlayedString = fmt.Sprintf("%d%d", col, row) + movesPlayedString
		}
		bb.TakeBack()
	}

	initialBoardString := bb.ToBoardString()

	for !tmpStack.isEmpty() {
		move := tmpStack.pop()
		bb.DoMove(move)
	}

	return initialBoardString + DELIMITER + movesPlayedString
}

func BoardStringToBitBoard(boardString string) BitBoard {
	var boardStringParts = strings.Split(boardString, ":")
	bbWhite, _ := strconv.ParseUint(boardStringParts[0], 10, 64)
	bbBlack, _ := strconv.ParseUint(boardStringParts[1], 10, 64)
	colorToMove, _ := strconv.Atoi(boardStringParts[2])
	var bitBoard = InitBoard(bbWhite, bbBlack, colorToMove)

	if len(boardStringParts) == 4 {
		var moveList = boardStringParts[3]
		for i := 0; i < len(moveList); i += 2 {
			var col = int(moveList[i] - '0')
			var row = int(moveList[i+1] - '0')
			if col == 8 && row == 8 {
				bitBoard.DoPassMove()
			} else {
				bitBoard.DoColRowMove(col, row)
			}
		}
	}
	return bitBoard
}
