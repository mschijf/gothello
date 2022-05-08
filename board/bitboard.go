package board

import (
	"gothello/bit64math"
	"gothello/collection"
)

type bitBoard struct {
	bitFields   [2]uint64
	colorToMove int
	stack       collection.Stack[move]
}

const rightBorder uint64 = 0x01_01_01_01_01_01_01_01
const leftBorder uint64 = 0x80_80_80_80_80_80_80_80
const verticalMiddle = ^(leftBorder | rightBorder)

const west = 1                  //1 shift to left
const northEast = BoardSize - 1 //7 shift to left
const north = BoardSize         //8 shift to left
const northWest = BoardSize + 1 //9 shift to left
const east = 1                  //1 shift to right
const southWest = BoardSize - 1 //7 shift to right
const south = BoardSize         //8 shift to right
const southEast = BoardSize + 1 //9 shift to right

const white = 0
const black = 1

func initBoard(bbWhite, bbBlack uint64, colorToMove int) bitBoard {
	var board = bitBoard{}
	board.bitFields[white] = bbWhite
	board.bitFields[black] = bbBlack
	board.colorToMove = colorToMove
	return board
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

func (bb *bitBoard) generateMoves() []move {
	var moveList []move

	var bbToMove = bb.bitFields[bb.colorToMove]
	var bbOpponent = bb.bitFields[1-bb.colorToMove]
	var bbEmpty = ^(bbToMove | bbOpponent)
	var bbWithoutLeftRightBorder = bbOpponent & verticalMiddle

	candWest := getLeftHittingCandidate(west, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candNorthEast := getLeftHittingCandidate(northEast, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candNorth := getLeftHittingCandidate(north, bbToMove, bbOpponent, bbEmpty)
	candNorthWest := getLeftHittingCandidate(northWest, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candEast := getRightHittingCandidate(east, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candSouthWest := getRightHittingCandidate(southWest, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candSouth := getRightHittingCandidate(south, bbToMove, bbOpponent, bbEmpty)
	candSouthEast := getRightHittingCandidate(southEast, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candAll := candWest | candNorthEast | candNorth | candNorthWest | candEast | candSouthWest | candSouth | candSouthEast

	for candAll != 0 {
		var allCaptures uint64 = 0
		var bbMove = bit64math.SmallesBit(candAll)
		if (bbMove & candWest) != 0 {
			allCaptures |= getLeftCapture(west, bbOpponent, bbMove)
		}
		if (bbMove & candNorthEast) != 0 {
			allCaptures |= getLeftCapture(northEast, bbOpponent, bbMove)
		}
		if (bbMove & candNorth) != 0 {
			allCaptures |= getLeftCapture(north, bbOpponent, bbMove)
		}
		if (bbMove & candNorthWest) != 0 {
			allCaptures |= getLeftCapture(northWest, bbOpponent, bbMove)
		}

		if (bbMove & candEast) != 0 {
			allCaptures |= getRightCapture(east, bbOpponent, bbMove)
		}
		if (bbMove & candSouthWest) != 0 {
			allCaptures |= getRightCapture(southWest, bbOpponent, bbMove)
		}
		if (bbMove & candSouth) != 0 {
			allCaptures |= getRightCapture(south, bbOpponent, bbMove)
		}
		if (bbMove & candSouthEast) != 0 {
			allCaptures |= getRightCapture(southEast, bbOpponent, bbMove)
		}

		moveList = append(moveList, move{allCaptures, bbMove})
		candAll ^= bbMove
	}
	if len(moveList) == 0 {
		return append(moveList, move{0, 0})
	}

	return moveList
}

func (bb *bitBoard) getAllCandidateMoves() uint64 {
	var bbToMove = bb.bitFields[bb.colorToMove]
	var bbOpponent = bb.bitFields[1-bb.colorToMove]
	var bbEmpty = ^(bbToMove | bbOpponent)
	var bbWithoutLeftRightBorder = bbOpponent & verticalMiddle

	candWest := getLeftHittingCandidate(west, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candNorthEast := getLeftHittingCandidate(northEast, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candNorth := getLeftHittingCandidate(north, bbToMove, bbOpponent, bbEmpty)
	candNorthWest := getLeftHittingCandidate(northWest, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candEast := getRightHittingCandidate(east, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candSouthWest := getRightHittingCandidate(southWest, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	candSouth := getRightHittingCandidate(south, bbToMove, bbOpponent, bbEmpty)
	candSouthEast := getRightHittingCandidate(southEast, bbToMove, bbWithoutLeftRightBorder, bbEmpty)
	return candWest | candNorthEast | candNorth | candNorthWest | candEast | candSouthWest | candSouth | candSouthEast
}

func (bb *bitBoard) doMove(move *move) {
	bb.bitFields[bb.colorToMove] ^= move.discsFlipped | move.discPlayed
	bb.colorToMove = 1 - bb.colorToMove
	bb.bitFields[bb.colorToMove] ^= move.discsFlipped
	bb.stack.Push(move)
}

func (bb *bitBoard) takeBack() {
	move := bb.stack.Pop()
	bb.bitFields[bb.colorToMove] ^= move.discsFlipped
	bb.colorToMove = 1 - bb.colorToMove
	bb.bitFields[bb.colorToMove] ^= move.discsFlipped | move.discPlayed
}

func (bb *bitBoard) isEndOfGame() bool {
	if ^(bb.bitFields[white] | bb.bitFields[black]) == 0 {
		return true
	}

	return bb.stack.Size() > 1 && bb.stack.FromTop(0).isPass() && bb.stack.FromTop(1).isPass()
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

func (bb *bitBoard) perft(depth int) int64 {
	if depth == 0 {
		return 1
	}
	if bb.isEndOfGame() {
		return 1
	}
	var nodeCount int64 = 0
	moves := bb.generateMoves()
	for _, move := range moves {
		bb.doMove(&move)
		nodeCount += bb.perft(depth - 1)
		bb.takeBack()
	}
	return nodeCount
}
