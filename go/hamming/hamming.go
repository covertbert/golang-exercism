package hamming

import (
	"errors"
	"strings"
)

// Distance calculates the number of variances between two strings
func Distance(a, b string) (int, error) {
	if len(b) != len(a) {
		return -1, errors.New("Error")
	}

	aArray := strings.Split(a, "")
	bArray := strings.Split(b, "")

	numberOfVariances := 0

	for index, letter := range aArray {
		if letter != bArray[index] {
			numberOfVariances++
		}
	}

	return numberOfVariances, nil
}
