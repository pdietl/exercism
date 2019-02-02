package collatzconjecture

import (
	"errors"
)

// CollatzConjecture returns the number of steps required through the
// CollatzConjecture to reduce the input to 1
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("input must be >= 0")
	}

	count := 0

	for ; n != 1; count++ {
		if n&1 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return count, nil
}
