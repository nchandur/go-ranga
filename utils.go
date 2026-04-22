package main

import "fmt"

// checks if file exists on board
func isValidFile(file File) bool {
	return file >= FileA && file <= FileH
}

// checks if rank exists on board
func isValidRank(rank Rank) bool {
	return rank >= Rank1 && rank <= Rank8
}

// returns 120-based index from file, rank
func FRTo120(file File, rank Rank) (Square, error) {

	if !isValidFile(file){
		return NoSquare, fmt.Errorf("%d: invalid file. must be between 1 and 8", file)
	}

	if !isValidRank(rank){
		return NoSquare, fmt.Errorf("%d: invalid rank. must be between 1 and 8", rank)
	}

	return Square((21 + int(file)) + (10 * int(rank))), nil
}
