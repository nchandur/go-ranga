package main

import "fmt"

// returns 120-based index from file, rank
func FRTo120(file File, rank Rank) (Square, error) {

	if !file.Valid() {
		return NoSquare, fmt.Errorf("%d: invalid file. must be between 1 and 8", file)
	}

	if !rank.Valid() {
		return NoSquare, fmt.Errorf("%d: invalid rank. must be between 1 and 8", rank)
	}

	return Square((21 + int(file)) + (10 * int(rank))), nil
}

// returns 64-based index from 120-based index
func SQ64(sq120 Square) (Square, error) {
	if sq120 < 0 || sq120 > 119 {
		return NoSquare, fmt.Errorf("invalid square value: %d", sq120)
	}

	return Sq120To64[sq120], nil
}

// returns 120-based index from 64-based index
func SQ120(sq64 Square) (Square, error) {
	if sq64 < 0 || sq64 > 63 {
		return NoSquare, fmt.Errorf("invalid square value: %d", sq64)
	}

	return Sq64To120[sq64], nil
}
