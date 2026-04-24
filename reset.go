package main

import "fmt"

func (b *Board) Reset() error {
	for idx := range BOARD_SQ_NUM {
		b.Pieces[idx] = Offboard
	}

	for idx := range 64 {
		sq64, err := SQ120(Square(idx))

		if err != nil {
			return fmt.Errorf("failed to reset board: %v", err)
		}

		b.Pieces[sq64] = Empty
	}

	b.BigPiece = [2]int{}
	b.MajorPiece = [2]int{}
	b.MinorPiece = [2]int{}
	b.Material = [2]int{}

	b.Pawns = [3]Bitboard{}
	b.PieceNum = [13]int{}

	b.KingsSquare[White] = NoSquare
	b.KingsSquare[Black] = NoSquare

	b.Side = Both
	b.EnPassant = NoSquare

	b.Ply = 0
	b.HistoryPly = 0
	b.CastleBit = 0

	b.PositionKey = uint64(0)
	b.History = [MAX_GAMES_MOVES]History{}

	return nil
}
