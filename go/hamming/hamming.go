// Package hamming computes the hamming distance between two DNA sequences
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance computes the hamming distance between two DNA sequences
// represented as two strings. The two sequences must be of equal length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("the two DNA sequences must be of the same length")
	}

	distance := 0

	for i, w := 0, 0; i < len(a); i += w {
		runValA, width := utf8.DecodeRuneInString(a[i:])
		runValB, _ := utf8.DecodeRuneInString(b[i:])
		w = width
		if runValA != runValB {
			distance++
		}
	}

	return distance, nil
}
