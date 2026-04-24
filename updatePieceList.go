package main

import "fmt"

func (b *Board) UpdatePieceList() error {

	for idx := range BOARD_SQ_NUM {
		sq := Square(idx)
		piece := b.Pieces[idx]

		if !piece.ValidEmptyOrOffBoard() {
			return fmt.Errorf("failed to update piece list. piece %s is invalid  at %s", &piece, &sq)
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
