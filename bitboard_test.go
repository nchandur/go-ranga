package main

import "testing"

func TestBitboardPop(t *testing.T) {
	tests := []struct {
		board    Bitboard
		expected int
		after    Bitboard
	}{
		{
			board:    0x1,
			expected: 0,
			after:    0x0,
		},
		{
			board:    0x3,
			expected: 0,
			after:    0x2,
		},
		{
			board:    0x100,
			expected: 8,
			after:    0x0,
		},
	}

	for _, test := range tests {
		tempBoard := test.board
		output := tempBoard.PopBit()
		if output != test.expected {
			t.Errorf("expected %d, want %d", output, test.expected)
		}

		if tempBoard != test.after {
			t.Errorf("expected %x, want %x", tempBoard, test.after)
		}
	}
}
func TestBitboardCount(t *testing.T) {
	tests := []struct {
		board    Bitboard
		expected int
	}{
		{board: 0x0001, expected: 1},
	}

	for _, test := range tests {
		output := test.board.Count()

		if output != test.expected {
			t.Errorf("expected: %d, output: %d", test.expected, output)
		}
	}

}

func TestBitboardSetBit(t *testing.T) {
	tests := []struct {
		board       Bitboard
		square      Square
		expected    Bitboard
		expectedErr bool
	}{
		{board: 0x0, square: A1, expected: 0x1, expectedErr: false},
		{board: 0x8, square: G1, expected: 0x48, expectedErr: false},
	}

	for _, test := range tests {

		err := test.board.SetBit(test.square)

		if err != nil && !test.expectedErr {
			t.Errorf("unexpected error: %v", err)
		}

		if test.board != test.expected {
			t.Errorf("expected %x, output: %x", test.expected, test.board)
		}

	}

}

func TestBitboardClearBit(t *testing.T) {
	tests := []struct {
		board       Bitboard
		square      Square
		expected    Bitboard
		expectedErr bool
	}{
		{board: 0x1, square: A1, expected: 0x0, expectedErr: false},
		{board: 0x8, square: G1, expected: 0x8, expectedErr: false},
		{board: 0xF, square: D1, expected: 0x7, expectedErr: false},
	}

	for _, test := range tests {

		err := test.board.ClearBit(test.square)

		if err != nil && !test.expectedErr {
			t.Errorf("unexpected error: %v", err)
		}

		if test.board != test.expected {
			t.Errorf("expected %x, output: %x", test.expected, test.board)
		}

	}

}
