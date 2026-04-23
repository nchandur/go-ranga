package main

import (
	"fmt"
	"strings"
)

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
	BigPiece   [2]int
	MajorPiece [2]int
	MinorPiece [2]int
	Material   [2]int

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

func NewBoard() (*Board, error) {
	board := Board{}

	if err := board.Reset(); err != nil {
		return nil, fmt.Errorf("failed to generate new board: %v", err)
	}

	return &board, nil

}

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

func (b *Board) UpdatePieceList() error {

	for idx := range BOARD_SQ_NUM {
		sq := Square(idx)
		piece := b.Pieces[idx]

		if !piece.ValidEmptyOrOffBoard() {
			return fmt.Errorf("failed to update piece list. piece %s is invalid or offboard at %s", &piece, &sq)
		}

		if piece != Offboard && piece != Empty {
			color := PieceColor[piece]

			if !color.Valid() {
				return fmt.Errorf("failed to update piece list. invalid color %s", &color)
			}

			if PieceBig[piece] {
				b.BigPiece[color]++
			}

			if PieceMajor[piece] {
				b.MajorPiece[color]++
			}

			if PieceMinor[piece] {
				b.MinorPiece[color]++
			}

			b.Material[color] += PieceValue[piece]

			if b.PieceNum[piece] > 10 || b.PieceNum[piece] < 0 {
				return fmt.Errorf("failed to update piece list. invalid number of piece %s on board. must be between 0 and 10. got %d", &piece, b.PieceNum[piece])
			}

			b.PieceList[piece][b.PieceNum[piece]] = idx
			b.PieceNum[piece]++

			if piece == wK {
				b.KingsSquare[White] = sq
			}

			if piece == bK {
				b.KingsSquare[Black] = sq
			}

			if piece == wP {
				if err := b.Pawns[White].SetBit(sq); err != nil {
					return fmt.Errorf("failed to update piece list: %v", err)
				}

				if err := b.Pawns[Both].SetBit(sq); err != nil {
					return fmt.Errorf("failed to update piece list: %v", err)
				}

			}

			if piece == bP {
				if err := b.Pawns[Black].SetBit(sq); err != nil {
					return fmt.Errorf("failed to update piece list: %v", err)
				}

				if err := b.Pawns[Both].SetBit(sq); err != nil {
					return fmt.Errorf("failed to update piece list: %v", err)
				}

			}

		}

	}

	return nil
}

func (b *Board) GeneratePositionKey() (uint64, error) {

	var res uint64

	for sq := range BOARD_SQ_NUM {
		piece := b.Pieces[sq]

		if piece != Empty && piece != Offboard {
			if !piece.Valid() {
				return uint64(0), fmt.Errorf("failed to generate position key. invalid piece: %s", &piece)
			}
			res ^= PieceKeys[piece][sq]
		}
	}

	if b.Side == White {
		res ^= SideKey
	}

	if b.EnPassant != NoSquare {
		if !b.EnPassant.Valid() {
			return uint64(0), fmt.Errorf("failed to generate position key. invalid en passant square %s", &b.EnPassant)
		}

		if !b.EnPassant.IsOnBoard() {
			return uint64(0), fmt.Errorf("failed to generate position key. en passant square %s is offboard", &b.EnPassant)
		}

		if FileBoard[b.EnPassant] != 3 && FileBoard[b.EnPassant] != 6 {
			return uint64(0), fmt.Errorf("failed to generate position key. en passant square must be on Rank 3 or 6, found on %d", FileBoard[b.EnPassant])
		}

		res ^= PieceKeys[Empty][b.EnPassant]

	}

	if !b.CastleBit.Valid() {
		return uint64(0), fmt.Errorf("failed to generate position key. invalid castle bit: %s", &b.CastleBit)
	}

	res ^= CastleKey[b.CastleBit]

	return res, nil
}

func (b *Board) String() string {
	var builder strings.Builder

	builder.WriteString("\n  +---------------------------------+\n")

	for rank := Rank8; rank >= Rank1; rank-- {
		fmt.Fprintf(&builder, "%d | ", rank+1)

		for file := FileA; file <= FileH; file++ {
			sq, _ := FRTo120(file, rank)
			piece := b.Pieces[sq]

			if piece == Empty {
				builder.WriteString(" .  ")
			} else {
				fmt.Fprintf(&builder, " %s  ", &piece)
			}
		}
		builder.WriteString("|\n")
	}

	builder.WriteString("  +---------------------------------+\n")

	builder.WriteString("     a   b   c   d   e   f   g   h\n\n")

	fmt.Fprintf(&builder, "Side To Move : %s\n", &b.Side)
	fmt.Fprintf(&builder, "En Passant   : %s\n", &b.EnPassant)
	fmt.Fprintf(&builder, "Castling     : %s\n", &b.CastleBit)
	fmt.Fprintf(&builder, "Pos Key      : %x\n", b.PositionKey)

	return builder.String()
}
