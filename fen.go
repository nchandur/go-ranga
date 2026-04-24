package main

import (
	"fmt"
	"strings"
)

const START string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func (b *Board) ParseFen(fen string) error {
	b.Reset()

	parts := strings.Fields(fen)
	if len(parts) < 4 {
		return fmt.Errorf("failed to parse FEN: expected at least 4 parts, got %d", len(parts))
	}

	rank := Rank8
	file := FileA

	for _, char := range parts[0] {
		count := 1
		var piece Piece

		switch char {
		case 'p':
			piece = bP
		case 'r':
			piece = bR
		case 'n':
			piece = bN
		case 'b':
			piece = bB
		case 'k':
			piece = bK
		case 'q':
			piece = bQ
		case 'P':
			piece = wP
		case 'R':
			piece = wR
		case 'N':
			piece = wN
		case 'B':
			piece = wB
		case 'K':
			piece = wK
		case 'Q':
			piece = wQ
		case '/':
			rank--
			file = FileA
			continue
		case '1', '2', '3', '4', '5', '6', '7', '8':
			piece = Empty
			count = int(char - '0')
		default:
			return fmt.Errorf("failed to parse FEN. invalid FEN character: %c", char)
		}

		for i := 0; i < count; i++ {
			sq64 := int(rank)*8 + int(file)
			sq120, err := SQ120(Square(sq64))
			if err != nil {
				return fmt.Errorf("failed to parse FEN: %v", err)
			}

			if piece != Empty {
				b.Pieces[sq120] = piece
			}
			file++
		}
	}

	if parts[1] == "w" {
		b.Side = White
	} else {
		b.Side = Black
	}

	for _, char := range parts[2] {
		switch char {
		case 'K':
			b.CastleBit |= WKCA
		case 'Q':
			b.CastleBit |= WQCA
		case 'k':
			b.CastleBit |= BKCA
		case 'q':
			b.CastleBit |= BQCA
		case '-':
		}
	}

	var err error

	if parts[3] != "-" {
		f := parts[3][0] - 'a'
		r := parts[3][1] - '1'
		b.EnPassant, err = FRTo120(File(f), Rank(r))

		if err != nil{
			return fmt.Errorf("failed to parse fen: %v", err)
		}

	}

	if b.PositionKey, err = b.GeneratePositionKey(); err != nil {
		return fmt.Errorf("failed to parse FEN: %v", err)
	}

	if err := b.UpdatePieceList(); err != nil {
		return fmt.Errorf("failed to parse FEN: %v", err)
	}

	return nil
}
