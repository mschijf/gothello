package pnsearch

type Node struct {
	pn, dpn   int
	isMaxNode bool
	childList []*Node
	bitFields [2]uint64
	//todo colorToMove
}

const Infinite = 999_999_999

//func (thisNode *Node) expand() {
//	bitBoard := board.InitBoard(thisNode.bitFields[0], thisNode.bitFields[1], 1)
//	moves := bitBoard.GenerateMoves()
//	thisNode.childList = make([]*Node, len(moves), len(moves))
//	for index, move := range moves {
//		bitBoard.DoMove(&move)
//		var pn, dpn int
//		if bitBoard.IsEndOfGame() {
//			pn, dpn = 0, Infinite
//		} else {
//			pn, dpn = 1, 1
//		}
//		thisNode.childList[index] = &Node{pn, dpn, !thisNode.isMaxNode, nil, bitBoard.GetBitFields()}
//		bitBoard.TakeBack()
//	}
//}
