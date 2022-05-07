package service

import "gothello/board"

type FieldModel struct {
	Col      int    `json:"col"`
	Row      int    `json:"row"`
	Color    string `json:"color"`
	Playable bool   `json:"playable"`
}

type BoardModel struct {
	Fields           [board.BoardSize][board.BoardSize]FieldModel `json:"fields"`
	ColorToMove      string                                       `json:"colorToMove"`
	TakeBackPossible bool                                         `json:"takeBackPossible"`
	GameFinished     bool                                         `json:"gameFinished"`
	MustPass         bool                                         `json:"mustPass"`
	BoardString      string                                       `json:"boardString"`
}

func ToBoardModel(humanBoard *board.HumanBoard) BoardModel {
	var bm = BoardModel{}
	for row := 0; row < board.BoardSize; row++ {
		for col := 0; col < board.BoardSize; col++ {
			bm.Fields[row][col] = getFieldModel(humanBoard, col, row)
		}
	}
	if humanBoard.IsBlackToMove() {
		bm.ColorToMove = "black"
	} else {
		bm.ColorToMove = "white"
	}
	bm.TakeBackPossible = humanBoard.HasHistory()
	bm.GameFinished = humanBoard.IsEndOfGame()
	bm.MustPass = humanBoard.MustPass() && !humanBoard.IsEndOfGame()
	bm.BoardString = humanBoard.ToBoardString()
	return bm
}

func getFieldModel(bb *board.HumanBoard, col, row int) FieldModel {
	var discColor string
	if bb.IsBlackDisc(col, row) {
		discColor = "black"
	} else if bb.IsWhiteDisc(col, row) {
		discColor = "white"
	} else {
		discColor = "NONE"
	}
	return FieldModel{col, row, discColor, bb.IsPlayable(col, row)}
}
