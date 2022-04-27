package board

import (
	"gothello/bit"
)

type BitBoard struct {
	Board       [2]uint64
	ColorToMove int
	stack       []Move
}

var RightBorder uint64 = 0x01_01_01_01_01_01_01_01
var LeftBorder uint64 = 0x80_80_80_80_80_80_80_80
var VerticalMiddle = ^(LeftBorder | RightBorder)

const WEST = 1      //1 shift to left
const NORTHEAST = 7 //7 shift to left
const NORTH = 8     //8 shift to left
const NORTHWEST = 9 //9 shift to left
const EAST = 1      //1 shift to right
const SOUTHWEST = 7 //7 shift to right
const SOUTH = 8     //8 shift to right
const SOUTHEAST = 9 //9 shift to right

const WHITE = 0
const BLACK = 1

func initBoard(bbWhite, bbBlack uint64, colorToMove int) BitBoard {
	var board = BitBoard{}
	board.Board[WHITE] = bbWhite
	board.Board[BLACK] = bbBlack
	board.ColorToMove = colorToMove
	return board
}

func InitStartBoard() BitBoard {
	return initBoard(0x00_00_00_10_08_00_00_00, 0x00_00_00_08_10_00_00_00, BLACK)
}

func getLeftHittingCandidate(direction int, bbToMove, bbCapturable, bbEmpty uint64) uint64 {
	var candidate uint64 = 0
	var loop = (bbToMove >> direction) & bbCapturable
	for loop != 0 {
		loop >>= direction
		candidate |= loop & bbEmpty
		loop &= bbCapturable
	}
	return candidate
}

func getLeftCapture(direction int, bbOpponent, bbMove uint64) uint64 {
	var allCaptures uint64 = 0
	var capture = bbMove << direction
	for ok := true; ok; ok = (capture & bbOpponent) != 0 {
		allCaptures |= capture
		capture <<= direction
	}
	return allCaptures
}

func getRightHittingCandidate(direction int, bbToMove, bbCapturable, bbEmpty uint64) uint64 {
	var candidate uint64 = 0
	var loop = (bbToMove << direction) & bbCapturable
	for loop != 0 {
		loop <<= direction
		candidate |= loop & bbEmpty
		loop &= bbCapturable
	}
	return candidate
}

func getRightCapture(direction int, bbOpponent, bbMove uint64) uint64 {
	var allCaptures uint64 = 0
	var capture = bbMove >> direction
	for ok := true; ok; ok = (capture & bbOpponent) != 0 {
		allCaptures |= capture
		capture >>= direction
	}
	return allCaptures
}

func (bb *BitBoard) GenerateMoves() []Move {
	var moveList []Move

	var bbToMove = bb.Board[bb.ColorToMove]
	var bbOpponent = bb.Board[1-bb.ColorToMove]
	var bbEmpty = ^(bbToMove | bbOpponent)
	var bbWithoutLeftRightBorder = bbOpponent & VerticalMiddle

	candWest := getLeftHittingCandidate(WEST, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candNorthEast := getLeftHittingCandidate(NORTHEAST, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candNorth := getLeftHittingCandidate(NORTH, bbToMove, bbOpponent, bbEmpty)
	candNorthWest := getLeftHittingCandidate(NORTHWEST, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candEast := getRightHittingCandidate(EAST, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candSouthWest := getRightHittingCandidate(SOUTHWEST, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candSouth := getRightHittingCandidate(SOUTH, bbToMove, bbOpponent, bbEmpty)
	candSouthEast := getRightHittingCandidate(SOUTHEAST, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candAll := candWest | candNorthEast | candNorth | candNorthWest | candEast | candSouthWest | candSouth | candSouthEast

	for candAll != 0 {
		var allCaptures uint64 = 0
		var bbMove = bit.SmallesBit(candAll)
		if (bbMove & candWest) != 0 {
			allCaptures |= getLeftCapture(WEST, bbOpponent, bbMove)
		}
		if (bbMove & candNorthEast) != 0 {
			allCaptures |= getLeftCapture(NORTHEAST, bbOpponent, bbMove)
		}
		if (bbMove & candNorth) != 0 {
			allCaptures |= getLeftCapture(NORTH, bbOpponent, bbMove)
		}
		if (bbMove & candNorthWest) != 0 {
			allCaptures |= getLeftCapture(NORTHWEST, bbOpponent, bbMove)
		}

		if (bbMove & candEast) != 0 {
			allCaptures |= getRightCapture(EAST, bbOpponent, bbMove)
		}
		if (bbMove & candSouthWest) != 0 {
			allCaptures |= getRightCapture(SOUTHWEST, bbOpponent, bbMove)
		}
		if (bbMove & candSouth) != 0 {
			allCaptures |= getRightCapture(SOUTH, bbOpponent, bbMove)
		}
		if (bbMove & candSouthEast) != 0 {
			allCaptures |= getRightCapture(SOUTHEAST, bbOpponent, bbMove)
		}

		moveList = append(moveList, Move{allCaptures, bbMove})
		candAll ^= bbMove
	}
	if len(moveList) == 0 {
		return append(moveList, Move{0, 0})
	}

	return moveList
}

func (bb *BitBoard) DoMove(move *Move) {
	bb.Board[bb.ColorToMove] ^= move.discsFlipped | move.discPlayed
	bb.ColorToMove = 1 - bb.ColorToMove
	bb.Board[bb.ColorToMove] ^= move.discsFlipped
	bb.stack = append(bb.stack, *move)
}

func (bb *BitBoard) TakeBack() {
	n := len(bb.stack) - 1 // Top element
	move := bb.stack[n]
	bb.stack = bb.stack[:n]

	bb.Board[bb.ColorToMove] ^= move.discsFlipped
	bb.ColorToMove = 1 - bb.ColorToMove
	bb.Board[bb.ColorToMove] ^= move.discsFlipped | move.discPlayed
}

func (bb *BitBoard) IsEndOfGame() bool {
	if ^(bb.Board[WHITE] | bb.Board[BLACK]) == 0 {
		return true
	}

	n := len(bb.stack) - 1
	return n > 1 && bb.stack[n].isPass() && bb.stack[n-1].isPass()
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//    depth   0  :     0.000000 ms -->              1
//    depth   1  :     0.000001 ms -->              4
//    depth   2  :     0.000001 ms -->             12
//    depth   3  :     0.000005 ms -->             56
//    depth   4  :     0.000017 ms -->            244
//    depth   5  :     0.000100 ms -->           1396
//    depth   6  :     0.000515 ms -->           8200
//    depth   7  :     0.003094 ms -->          55092
//    depth   8  :     0.021635 ms -->         390216
//    depth   9  :     0.137550 ms -->        3005288
//    depth  10  :     1.109636 ms -->       24571284
//    depth  11  :     9.388059 ms -->      212258800
//    depth  12  :    83.742958 ms -->     1939886636
//    depth  13  :   782.551742 ms -->    18429641748
//
//    speed: 23.550.700 per second
//
//    see also http://www.aartbik.com/strategy.php
//
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (bb *BitBoard) Perft(depth int) int64 {
	if depth == 0 {
		return 1
	}
	if bb.IsEndOfGame() {
		return 1
	}
	var nodeCount int64 = 0
	moves := bb.GenerateMoves()
	for _, move := range moves {
		bb.DoMove(&move)
		nodeCount += bb.Perft(depth - 1)
		bb.TakeBack()
	}
	return nodeCount
}
