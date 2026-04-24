package main

import "fmt"

func (b *Board) IsAttacked(square Square, side Color) (bool, error) {

	if !square.IsOnBoard() {
		return false, fmt.Errorf("square %s is offboard", &square)
	}

	if !side.Valid() {
		return false, fmt.Errorf("side %s is invalid", &side)
	}

	if _, err := b.Check(); err != nil {
		return false, err
	}

	ok, err := b.isAttackedByPawn(square, side)

	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	ok, err = b.isAttackedByKnight(square, side)

	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	ok, err = b.isAttackedByQueenOrRook(square, side)

	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	ok, err = b.isAttackedByQueenOrBishop(square, side)

	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	ok, err = b.isAttackedByKing(square, side)

	if err != nil {
		return false, err
	}

	if ok {
		return true, nil
	}

	return false, nil
}

func (b *Board) isAttackedByPawn(square Square, side Color) (bool, error) {

	if side == White {
		if b.Pieces[square-11] == wP || b.Pieces[square-9] == wP {
			return true, nil
		}
	}

	if side == Black {
		if b.Pieces[square+11] == bP || b.Pieces[square+9] == bP {
			return true, nil
		}
	}

	return false, nil
}

func (b *Board) isAttackedByKnight(square Square, side Color) (bool, error) {

	for idx := range 8 {
		pce := b.Pieces[square+Square(KnightDirections[idx])]

		if !pce.ValidEmptyOrOffBoard() {
			return false, fmt.Errorf("square is neither valid, empty or offboard")
		}

		if pce != Offboard && IsKnight[pce] && PieceColor[pce] == side {
			return true, nil
		}

	}

	return false, nil
}

func (b *Board) isAttackedByQueenOrRook(square Square, side Color) (bool, error) {

	for idx := range 4 {
		dir := RookDirections[idx]

		tSq := square + Square(dir)

		if tSq < 0 || tSq >= 120 {
			return false, fmt.Errorf("square index not valid. must be between 0 and 120.")
		}

		pce := b.Pieces[tSq]

		if !pce.ValidEmptyOrOffBoard() {
			return false, fmt.Errorf("invalid piece")
		}

		for pce != Offboard {
			if pce != Empty {
				if (IsRook[pce] || IsQueen[pce]) && PieceColor[pce] == side {
					return true, nil
				}
				break
			}

			tSq += Square(dir)

			if tSq < 0 || tSq >= 120 {
				return false, fmt.Errorf("square index not valid. must be between 0 and 120.")
			}

			pce = b.Pieces[tSq]

		}

	}

	return false, nil

}

func (b *Board) isAttackedByQueenOrBishop(square Square, side Color) (bool, error) {
	for idx := range 4 {
		dir := BishopDirections[idx]

		tSq := square + Square(dir)

		if tSq < 0 || tSq >= 120 {
			return false, fmt.Errorf("square index not valid. must be between 0 and 120.")
		}

		pce := b.Pieces[tSq]

		if !pce.ValidEmptyOrOffBoard() {
			return false, fmt.Errorf("invalid piece")
		}

		for pce != Offboard {
			if pce != Empty {
				if (IsBishop[pce] || IsQueen[pce]) && PieceColor[pce] == side {
					return true, nil
				}
				break
			}

			tSq += Square(dir)

			if tSq < 0 || tSq >= 120 {
				return false, fmt.Errorf("square index not valid. must be between 0 and 120.")
			}

			pce = b.Pieces[tSq]

		}

	}

	return false, nil

}

func (b *Board) isAttackedByKing(square Square, side Color) (bool, error) {

	for idx := range 8 {
		pce := b.Pieces[square+Square(KingDirections[idx])]

		if !pce.ValidEmptyOrOffBoard() {
			return false, fmt.Errorf("square is neither valid, empty or offboard")
		}

		if pce != Offboard && IsKing[pce] && PieceColor[pce] == side {
			return true, nil
		}

	}

	return false, nil

}
