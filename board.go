package main

type Board struct {
	Pieces [BOARD_SQ_NUM]Square
	Pawns  [3]uint64

	KingsSquare [2]Square

	Side      Color
	EnPassant Square

	FiftyMove  int
	Ply        int
	HistoryPly int

	PositionKey uint64

	PieceNum   [13]Piece
	BigPiece   [3]int
	MajorPiece [3]int
	MinorPiece [3]int

	History [MAX_GAMES_MOVES]History
}

type History struct {
	Move int
	CastleBit
	EnPassant   Square
	FiftyMove   int
	PositionKey uint64
}
