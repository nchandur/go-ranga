package main

import "fmt"

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

		if RankBoard[b.EnPassant] != Rank3 && RankBoard[b.EnPassant] != Rank6 {
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
