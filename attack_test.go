package main

import "testing"

func TestIsAttackedByPawn(t *testing.T) {
	tests := []struct {
		fen      string
		square   Square
		side     Color
		expected bool
	}{
		{fen: "3k4/3p4/8/8/8/8/8/3K4 b - - 0 1", square: C6, side: Black, expected: true},
		{fen: "3k4/3p4/8/8/8/8/8/3K4 b - - 0 1", square: E6, side: Black, expected: true},
		{fen: "3k4/3p4/8/8/8/8/8/3K4 b - - 0 1", square: D6, side: Black, expected: false},

		{fen: "3k4/8/8/8/8/8/3P4/3K4 w - - 0 1", square: C3, side: White, expected: true},
		{fen: "3k4/8/8/8/8/8/3P4/3K4 w - - 0 1", square: E3, side: White, expected: true},
		{fen: "3k4/8/8/8/8/8/3P4/3K4 w - - 0 1", square: D3, side: White, expected: false},
	}

	for _, test := range tests {

		board, err := NewBoard()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		err = board.ParseFen(test.fen)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		_, err = board.Check()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		output, err := board.isAttackedByPawn(test.square, test.side)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}
	}

}

func TestIsAttackedByKnight(t *testing.T) {
	tests := []struct {
		fen      string
		square   Square
		side     Color
		expected bool
	}{
		{fen: "3k4/8/8/4n3/8/8/8/3K4 b - - 0 1", square: D7, side: Black, expected: true},
		{fen: "3k4/8/8/4n3/8/8/8/3K4 b - - 0 1", square: F7, side: Black, expected: true},
		{fen: "3k4/8/8/4n3/8/8/8/3K4 b - - 0 1", square: C6, side: Black, expected: true},
		{fen: "3k4/8/8/4n3/8/8/8/3K4 b - - 0 1", square: G6, side: Black, expected: true},
		{fen: "3k4/8/8/4n3/8/8/8/3K4 b - - 0 1", square: C4, side: Black, expected: true},
		{fen: "3k4/8/8/4n3/8/8/8/3K4 b - - 0 1", square: G4, side: Black, expected: true},
		{fen: "3k4/8/8/4n3/8/8/8/3K4 b - - 0 1", square: D3, side: Black, expected: true},
		{fen: "3k4/8/8/4n3/8/8/8/3K4 b - - 0 1", square: F3, side: Black, expected: true},
	}

	for _, test := range tests {

		board, err := NewBoard()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		err = board.ParseFen(test.fen)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		_, err = board.Check()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		output, err := board.isAttackedByKnight(test.square, test.side)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}
	}

}

func TestIsAttackedByQueenOrRook(t *testing.T) {
	tests := []struct {
		fen      string
		square   Square
		side     Color
		expected bool
	}{
		{fen: "3k4/8/8/4r3/8/8/8/3K4 b - - 0 1", square: F5, side: Black, expected: true},
		{fen: "3k4/8/8/4r3/8/8/8/3K4 b - - 0 1", square: D5, side: Black, expected: true},
		{fen: "3k4/8/8/4r3/8/8/8/3K4 b - - 0 1", square: E7, side: Black, expected: true},
		{fen: "3k4/8/8/4r3/8/8/8/3K4 b - - 0 1", square: E3, side: Black, expected: true},
		{fen: "3k4/8/8/4q3/8/8/8/3K4 b - - 0 1", square: E8, side: Black, expected: true},
		{fen: "3k4/8/8/4q3/8/8/8/3K4 b - - 0 1", square: E1, side: Black, expected: true},
		{fen: "3k4/8/8/4q3/8/8/8/3K4 b - - 0 1", square: A5, side: Black, expected: true},
		{fen: "3k4/8/8/4q3/8/8/8/3K4 b - - 0 1", square: H5, side: Black, expected: true},
	}

	for _, test := range tests {

		board, err := NewBoard()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		err = board.ParseFen(test.fen)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		_, err = board.Check()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		output, err := board.isAttackedByQueenOrRook(test.square, test.side)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}
	}
}

func TestIsAttackedByQueenOrBishop(t *testing.T) {
	tests := []struct {
		fen      string
		square   Square
		side     Color
		expected bool
	}{
		{fen: "3k4/8/8/4b3/8/8/8/3K4 b - - 0 1", square: D6, side: Black, expected: true},
		{fen: "3k4/8/8/4b3/8/8/8/3K4 b - - 0 1", square: F4, side: Black, expected: true},
		{fen: "3k4/8/8/4b3/8/8/8/3K4 b - - 0 1", square: F6, side: Black, expected: true},
		{fen: "3k4/8/8/4b3/8/8/8/3K4 b - - 0 1", square: D4, side: Black, expected: true},
		{fen: "3k4/8/8/4q3/8/8/8/3K4 b - - 0 1", square: C7, side: Black, expected: true},
		{fen: "3k4/8/8/4q3/8/8/8/3K4 b - - 0 1", square: B8, side: Black, expected: true},
		{fen: "3k4/8/8/4q3/8/8/8/3K4 b - - 0 1", square: G3, side: Black, expected: true},
		{fen: "3k4/8/8/4q3/8/8/8/3K4 b - - 0 1", square: H2, side: Black, expected: true},
	}

	for _, test := range tests {

		board, err := NewBoard()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		err = board.ParseFen(test.fen)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		_, err = board.Check()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		output, err := board.isAttackedByQueenOrBishop(test.square, test.side)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}
	}
}

func TestIsAttackedByKing(t *testing.T) {
	tests := []struct {
		fen      string
		square   Square
		side     Color
		expected bool
	}{
		{fen: "8/8/8/4k3/8/8/8/3K4 b - - 0 1", square: E6, side: Black, expected: true},
		{fen: "8/8/8/4k3/8/8/8/3K4 b - - 0 1", square: E4, side: Black, expected: true},
		{fen: "8/8/8/4k3/8/8/8/3K4 b - - 0 1", square: D5, side: Black, expected: true},
		{fen: "8/8/8/4k3/8/8/8/3K4 b - - 0 1", square: F5, side: Black, expected: true},
		{fen: "8/8/8/4k3/8/8/8/3K4 b - - 0 1", square: F6, side: Black, expected: true},
		{fen: "8/8/8/4k3/8/8/8/3K4 b - - 0 1", square: F4, side: Black, expected: true},
		{fen: "8/8/8/4k3/8/8/8/3K4 b - - 0 1", square: D6, side: Black, expected: true},
		{fen: "8/8/8/4k3/8/8/8/3K4 b - - 0 1", square: D4, side: Black, expected: true},
	}

	for _, test := range tests {

		board, err := NewBoard()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		err = board.ParseFen(test.fen)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		_, err = board.Check()

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		output, err := board.isAttackedByKing(test.square, test.side)

		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}
	}

}
