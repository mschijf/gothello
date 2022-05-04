package service

import (
	"gothello/board"
)

func GetNewBoard() (BoardModel, string) {
	initialBoard := board.InitStartBoard()
	return BitBoardToBoardModel(&initialBoard), initialBoard.ToBoardStatusString()
}

func GetBoard(boardStatusString string) (BoardModel, string) {
	currentBoard := board.BoardStringToBitBoard(boardStatusString)
	return BitBoardToBoardModel(&currentBoard), currentBoard.ToBoardStatusString()
}

func DoMove(boardStatusString string, col, row int) (BoardModel, string) {
	currentBoard := board.BoardStringToBitBoard(boardStatusString)
	currentBoard.DoColRowMove(col, row)
	return BitBoardToBoardModel(&currentBoard), currentBoard.ToBoardStatusString()
}

func DoPassMove(boardStatusString string) (BoardModel, string) {
	currentBoard := board.BoardStringToBitBoard(boardStatusString)
	currentBoard.DoPassMove()
	return BitBoardToBoardModel(&currentBoard), currentBoard.ToBoardStatusString()
}

func TakeBackLastMove(boardStatusString string) (BoardModel, string) {
	currentBoard := board.BoardStringToBitBoard(boardStatusString)
	currentBoard.TakeBack()
	return BitBoardToBoardModel(&currentBoard), currentBoard.ToBoardStatusString()
}
