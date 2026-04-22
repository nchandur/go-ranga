package main

import (
	"testing"
)

func TestIsValidFile(t *testing.T) {
	tests := []struct {
		file     File
		expected bool
	}{
		{file: 1, expected: true},
		{file: 7, expected: true},
		{file: -1, expected: false},
		{file: 10, expected: false},
	}

	for _, test := range tests {
		output := isValidFile(test.file)

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}

	}

}

func TestIsValidRank(t *testing.T) {
	tests := []struct {
		rank     Rank
		expected bool
	}{
		{rank: 1, expected: true},
		{rank: 7, expected: true},
		{rank: -1, expected: false},
		{rank: 10, expected: false},
	}

	for _, test := range tests {
		output := isValidRank(test.rank)

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}

	}

}

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
