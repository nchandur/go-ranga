package main

import "fmt"

func (b *Board) Check() (bool, error) {

	tempPieceNum := [13]int{}
	tempBigPiece := [2]int{}
	tempMajorPiece := [2]int{}
	tempMinorPiece := [2]int{}
	tempMaterial := [2]int{}

	tempPawns := [3]Bitboard{}

	tempPawns[White] = b.Pawns[White]
	tempPawns[Black] = b.Pawns[Black]
	tempPawns[Both] = b.Pawns[Both]

	for t := wP; t <= bK; t++ {
		for n := range b.PieceNum[t] {
			sq := Square(b.PieceList[t][n])

			if b.Pieces[sq] != t {
				return false, fmt.Errorf("pieces %s %s do not match on square %s", &b.Pieces[sq], &t, &sq)
			}
		}
	}

	for sq := range 64 {
		sq120, _ := SQ120(Square(sq))
		tpce := b.Pieces[sq120]
		tempPieceNum[tpce]++
		color := PieceColor[tpce]

		if color != Both {
			if PieceBig[tpce] {
				tempBigPiece[color]++
			}

			if PieceMajor[tpce] {
				tempMajorPiece[color]++
			}

			if PieceMinor[tpce] {
				tempMinorPiece[color]++
			}

			tempMaterial[color] += PieceValue[tpce]
		}

	}

	for tpce := wP; tpce <= bK; tpce++ {
		if tempPieceNum[tpce] != b.PieceNum[tpce] {
			return false, fmt.Errorf("piece numbers do not match for piece %s. %d != %d", &tpce, b.PieceNum[tpce], tempPieceNum[tpce])
		}
	}

	pcount := tempPawns[White].Count()

	if pcount != b.PieceNum[wP] {
		return false, fmt.Errorf("white pawn count does not match %d != %d", pcount, b.PieceNum[wP])
	}

	pcount = tempPawns[Black].Count()

	if pcount != b.PieceNum[bP] {
		return false, fmt.Errorf("black pawn count does not match %d != %d", pcount, b.PieceNum[bP])
	}

	pcount = tempPawns[Both].Count()

	if pcount != b.PieceNum[wP]+b.PieceNum[bP] {
		return false, fmt.Errorf("pawn count does not match %d != %d", pcount, b.PieceNum[wP]+b.PieceNum[bP])
	}

	for tempPawns[White] != 0 {
		sq64 := tempPawns[White].PopBit()
		sq, err := SQ120(Square(sq64))

		if err != nil {
			fmt.Println(&tempPawns[White])
			return false, fmt.Errorf("failed to update piece list: %v", err)
		}

		if b.Pieces[sq] != wP {
			return false, fmt.Errorf("%s must contain white pawn", &sq)
		}
	}

	for tempPawns[Black] != 0 {
		sq64 := tempPawns[Black].PopBit()
		sq, err := SQ120(Square(sq64))

		if err != nil {
			return false, fmt.Errorf("failed to update piece list: %v", err)
		}

		if b.Pieces[sq] != bP {
			return false, fmt.Errorf("%s must contain black pawn", &sq)
		}
	}

	for tempPawns[Both] != 0 {
		sq64 := tempPawns[Both].PopBit()
		sq, err := SQ120(Square(sq64))

		if err != nil {
			return false, fmt.Errorf("failed to update piece list: %v", err)
		}

		if b.Pieces[sq] != bP && b.Pieces[sq] != wP {
			return false, fmt.Errorf("%s must contain pawn", &sq)
		}
	}

	if tempMaterial[White] != b.Material[White] {
		return false, fmt.Errorf("white material does not match. %d != %d", tempMaterial[White], b.Material[White])
	}

	if tempMaterial[Black] != b.Material[Black] {
		return false, fmt.Errorf("black material does not match. %d != %d", tempMaterial[Black], b.Material[Black])
	}

	if tempMinorPiece[White] != b.MinorPiece[White] {
		return false, fmt.Errorf("white minor pieces do not match. %d != %d", tempMinorPiece[White], b.MinorPiece[White])
	}

	if tempMinorPiece[Black] != b.MinorPiece[Black] {
		return false, fmt.Errorf("black minor pieces do not match. %d != %d", tempMinorPiece[Black], b.MinorPiece[Black])
	}

	if tempMajorPiece[White] != b.MajorPiece[White] {
		return false, fmt.Errorf("white major pieces do not match. %d != %d", tempMajorPiece[White], b.MajorPiece[White])
	}

	if tempMajorPiece[Black] != b.MajorPiece[Black] {
		return false, fmt.Errorf("black major pieces do not match. %d != %d", tempMajorPiece[Black], b.MajorPiece[Black])
	}

	if tempBigPiece[White] != b.BigPiece[White] {
		return false, fmt.Errorf("white big pieces do not match. %d != %d", tempBigPiece[White], b.BigPiece[White])
	}

	if tempBigPiece[Black] != b.BigPiece[Black] {
		return false, fmt.Errorf("black big pieces do not match. %d != %d", tempBigPiece[Black], b.BigPiece[Black])
	}

	if b.Side != White && b.Side != Black {
		return false, fmt.Errorf("side to play invalid: %s", &b.Side)
	}

	tempKey, err := b.GeneratePositionKey()

	if err != nil {
		return false, err
	}

	if tempKey != b.PositionKey {
		return false, fmt.Errorf("position keys do not match. %x != %x", tempKey, b.PositionKey)
	}

	if !(b.EnPassant == NoSquare || (RankBoard[b.EnPassant] == Rank6 && b.Side == White) || (RankBoard[b.EnPassant] == Rank3 && b.Side != Black)) {
		return false, fmt.Errorf("invalid enpassant square %s", &b.EnPassant)
	}

	if b.Pieces[b.KingsSquare[White]] != wK {
		return false, fmt.Errorf("white king must be on square %s", &b.KingsSquare[White])
	}

	if b.Pieces[b.KingsSquare[Black]] != bK {
		return false, fmt.Errorf("black king must be on square %s", &b.KingsSquare[Black])
	}

	if !b.CastleBit.Valid() {
		return false, fmt.Errorf("invalid castle bit %s", &b.CastleBit)
	}

	// 	ASSERT(PceListOk(pos));

	return true, nil
}
