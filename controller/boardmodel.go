package controller

//album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//const val MAX_ROW = 8
//const val MAX_COL = 8
//
//class BoardModel(board: HumanBoard)  {
//
//val fields = Array(MAX_ROW) { row -> Array(MAX_COL) { col -> getFieldModel(board, col, row) } }
//val colorToMove = if (board.isBlackToMove()) Color.BLACK else Color.WHITE
//val takeBackPossible = board.hasHistory()
//val gameFinished = board.isEndOfGame()
//val mustPass = board.mustPass() && !board.isEndOfGame()
//val boardString = board.toBoardString()
//
//private fun getFieldModel(board: HumanBoard, col: Int, row: Int): FieldModel {
//return FieldModel(
//col=col,
//row=row,
//color=if (board.isWhiteDisc(col, row)) Color.WHITE else if (board.isBlackDisc(col, row)) Color.BLACK else Color.NONE,
//playable=board.isPlayable(col,row))
//}
//}
//
//data class FieldModel(val col: Int, val row: Int, val color: Color, val playable: Boolean)
