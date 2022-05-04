package service

import "gothello/board"

const MAX_ROW = 8
const MAX_COL = 8

type FieldModel struct {
	Col      int    `json:"col"`
	Row      int    `json:"row"`
	Color    string `json:"color"`
	Playable bool   `json:"playable"`
}

type BoardModel struct {
	Fields           [MAX_ROW][MAX_COL]FieldModel `json:"fields"`
	ColorToMove      string                       `json:"colorToMove"`
	TakeBackPossible bool                         `json:"takeBackPossible"`
	GameFinished     bool                         `json:"gameFinished"`
	MustPass         bool                         `json:"mustPass"`
	BoardString      string                       `json:"boardString"`
}

func BitBoardToBoardModel(bb *board.BitBoard) BoardModel {
	var bm = BoardModel{}
	for row := 0; row < MAX_ROW; row++ {
		for col := 0; col < MAX_COL; col++ {
			bm.Fields[row][col] = getFieldModel(bb, col, row)
		}
	}
	if bb.IsBlackToMove() {
		bm.ColorToMove = "BLACK"
	} else {
		bm.ColorToMove = "WHITE"
	}
	bm.TakeBackPossible = bb.HasHistory()
	bm.GameFinished = bb.IsEndOfGame()
	bm.MustPass = bb.MustPass() && !bb.IsEndOfGame()
	bm.BoardString = bb.ToBoardString()
	return bm
}

func getFieldModel(bb *board.BitBoard, col, row int) FieldModel {
	var discColor string
	if bb.IsBlackDisc(col, row) {
		discColor = "BLACK"
	} else if bb.IsWhiteDisc(col, row) {
		discColor = "WHITE"
	} else {
		discColor = "NONE"
	}
	return FieldModel{col, row, discColor, bb.IsPlayable(col, row)}
}
