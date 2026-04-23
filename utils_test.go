package main

import (
	"testing"
)

func TestFRTo120(t *testing.T) {
	tests := []struct {
		file        File
		rank        Rank
		expected    Square
		expectedErr bool
	}{
		{file: 0, rank: 0, expected: 21, expectedErr: false},
		{file: 1, rank: 1, expected: 32, expectedErr: false},
		{file: 0, rank: 5, expected: 71, expectedErr: false},
		{file: 10, rank: -1, expected: NoSquare, expectedErr: true},
	}

	for _, test := range tests {

		output, err := FRTo120(test.file, test.rank)

		if err != nil && !test.expectedErr {
			t.Errorf("unexpected error: %v", err)
		}

		if output != test.expected {
			t.Errorf("expected %d, output: %d", test.expected, output)
		}

	}

}

func TestSq120(t *testing.T) {

	tests := []struct {
		sq64        Square
		expected    Square
		expectedErr bool
	}{
		{sq64: 0, expected: 21, expectedErr: false},
		{sq64: 63, expected: 98, expectedErr: false},
		{sq64: 20, expected: 45, expectedErr: true},
		{sq64: 100, expected: NoSquare, expectedErr: true},
	}

	for _, test := range tests {

		output, err := SQ120(test.sq64)

		if err != nil && !test.expectedErr {
			t.Errorf("unexpected error: %v", err)
		}

		if output != test.expected {
			t.Errorf("expected %d, output: %d", test.expected, output)
		}

	}

}

func TestSq64(t *testing.T) {

	tests := []struct {
		sq120        Square
		expected    Square
		expectedErr bool
	}{
		{sq120: 0, expected: 65, expectedErr: false},
		{sq120: 63, expected: 34, expectedErr: false},
		{sq120: 22, expected: 1, expectedErr: false},
		{sq120: 120, expected: NoSquare, expectedErr: true},
	}

	for _, test := range tests {

		output, err := SQ64(test.sq120)

		if err != nil && !test.expectedErr {
			t.Errorf("unexpected error: %v", err)
		}

		if output != test.expected {
			t.Errorf("input: %d, expected %d, output: %d", test.sq120, test.expected, output)
		}

	}

}
