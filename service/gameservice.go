package service

import (
	"gothello/board"
)

var currentBoard = board.InitStartBoard()

func GetNewBoard() BoardModel {
	currentBoard = board.InitStartBoard()
	return BitBoardToBoardModel(&currentBoard)
}

func GetBoard() BoardModel {
	return BitBoardToBoardModel(&currentBoard)
}

func DoMove(col, row int) BoardModel {
	currentBoard.DoColRowMove(col, row)
	return BitBoardToBoardModel(&currentBoard)
}

func DoPassMove() BoardModel {
	currentBoard.DoPassMove()
	return BitBoardToBoardModel(&currentBoard)
}

func TakeBackLastMove() BoardModel {
	currentBoard.TakeBack()
	return BitBoardToBoardModel(&currentBoard)
}
