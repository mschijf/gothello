package board

import (
	"gothello/math/bit64math"
)

type BitBoard [2]uint64

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

func InitBoard(bbWhite, bbBlack uint64) BitBoard {
	return BitBoard{bbWhite, bbBlack}
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

func (bitBoard *BitBoard) GeneratePositions(colorToMove int) []BitBoard {
	var resultList []BitBoard

	var opponentColor = 1 - colorToMove
	var bbToMove = bitBoard[colorToMove]
	var bbOpponent = bitBoard[opponentColor]
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

		var bitFields [2]uint64
		bitFields[colorToMove] = bbToMove ^ (allCaptures | bbMove)
		bitFields[opponentColor] = bbOpponent ^ allCaptures

		resultList = append(resultList, bitFields)
		candAll ^= bbMove
	}
	return resultList
}

func (bitBoard *BitBoard) GenerateMoves(colorToMove int) []Move {
	var resultList []Move

	var opponentColor = 1 - colorToMove
	var bbToMove = bitBoard[colorToMove]
	var bbOpponent = bitBoard[opponentColor]
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

		resultList = append(resultList, Move{allCaptures, bbMove})
		candAll ^= bbMove
	}
	if len(resultList) == 0 {
		resultList = append(resultList, Move{0, 0})
	}
	return resultList
}

func (bitBoard *BitBoard) getAllCandidateMoves(colorToMove int) uint64 {
	var bbToMove = bitBoard[colorToMove]
	var bbOpponent = bitBoard[1-colorToMove]
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

func (bitBoard *BitBoard) DoMove(move *Move, movingColor int) {
	bitBoard[movingColor] ^= move.discsFlipped | move.discPlayed
	bitBoard[1-movingColor] ^= move.discsFlipped
}

func (bitBoard *BitBoard) UndoMove(move *Move, movingColor int) {
	bitBoard[1-movingColor] ^= move.discsFlipped
	bitBoard[movingColor] ^= move.discsFlipped | move.discPlayed
}

func (bitBoard *BitBoard) AllFieldsPlayed() bool {
	return ^(bitBoard[white] | bitBoard[black]) == 0
}

func (bitBoard *BitBoard) ColorHasWon(color int) bool {
	return bit64math.BitCount(bitBoard[color]) > bit64math.BitCount(bitBoard[1-color])
}
