package main

import "fmt"

type Board struct {
	Pieces [BOARD_SQ_NUM]Piece
	Pawns  [3]Bitboard

	KingsSquare [2]Square

	Side      Color
	EnPassant Square

	FiftyMove  int
	Ply        int
	HistoryPly int

	PositionKey uint64

	CastleBit
	PieceNum   [13]int
	BigPiece   [3]int
	MajorPiece [3]int
	MinorPiece [3]int

	History [MAX_GAMES_MOVES]History

	// 13: piece number each side
	// 10: squares on which pieces exist
	PieceList [13][10]int
}

type History struct {
	Move int
	CastleBit
	EnPassant   Square
	FiftyMove   int
	PositionKey uint64
}

func NewBoard() Board {
	board := Board{}

	for idx := range board.KingsSquare {
		board.KingsSquare[idx] = NoSquare
	}

	board.EnPassant = NoSquare

	return board

}

func (b *Board) GeneratePositionKey() (uint64, error) {

	var res uint64

	for sq := range BOARD_SQ_NUM {
		piece := b.Pieces[sq]

		if piece != Empty {
			if !piece.Valid() {
				return uint64(0), fmt.Errorf("failed to generate position key. invalid piece: %d", piece)
			}
			res ^= PieceKeys[piece][sq]
		}
	}

	if b.Side == White {
		res ^= SideKey
	}

	if b.EnPassant != NoSquare {
		if !b.EnPassant.Valid() {
			return uint64(0), fmt.Errorf("failed to generate position key. invalid en passant square %d", b.EnPassant)
		}

		if !b.EnPassant.IsOnBoard() {
			return uint64(0), fmt.Errorf("failed to generate position key. en passant square %d is offboard", b.EnPassant)
		}

		if FileBoard[b.EnPassant] != 3 && FileBoard[b.EnPassant] != 6 {
			return uint64(0), fmt.Errorf("failed to generate position key. en passant square must be on Rank 3 or 6, found on %d", FileBoard[b.EnPassant])
		}

		res ^= PieceKeys[Empty][b.EnPassant]

	}

	if !b.CastleBit.Valid() {
		return uint64(0), fmt.Errorf("failed to generate position key. invalid castle bit: %d", b.CastleBit)
	}

	res ^= CastleKey[b.CastleBit]

	return res, nil
}
