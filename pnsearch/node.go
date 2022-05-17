package pnsearch

import (
	"gothello/board"
)

type Node struct {
	pn, dpn   int
	isMaxNode bool
	parent    *Node
	childList []*Node
	bitBoard  board.BitBoard
}

const Infinite = 999_999_999

func initPnDpn(position board.BitBoard, endPosition bool, maxNodeColor int) (pn int, dpn int) {
	if !endPosition {
		return 1, 1
	}

	if position.ColorHasWon(maxNodeColor) {
		return 0, Infinite
	}
	return Infinite, 0
}

func (thisNode *Node) expand(colorToMove int, maxNodeColor int) {
	newPositions := thisNode.bitBoard.GeneratePositions(colorToMove)
	if len(newPositions) == 0 {
		doublePass := thisNode.bitBoard == thisNode.parent.bitBoard
		newPosition := thisNode.bitBoard
		pn, dpn := initPnDpn(newPosition, doublePass, maxNodeColor)
		thisNode.childList = make([]*Node, 1, 1)
		thisNode.childList[0] = &Node{pn, dpn, !thisNode.isMaxNode, thisNode, nil, newPosition}
	} else {
		thisNode.childList = make([]*Node, len(newPositions), len(newPositions))
		for index, newPosition := range newPositions {
			pn, dpn := initPnDpn(newPosition, newPosition.AllFieldsPlayed(), maxNodeColor)
			thisNode.childList[index] = &Node{pn, dpn, !thisNode.isMaxNode, thisNode, nil, newPosition}
		}
	}
}
