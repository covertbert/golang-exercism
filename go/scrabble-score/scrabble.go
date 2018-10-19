package scrabble

import (
	"strings"
)

var scoreMap = map[string]int{"A": 1, "B": 3, "C": 3, "D": 2, "E": 1, "F": 4, "G": 2, "H": 4, "I": 1, "J": 8, "K": 5, "L": 1, "M": 3, "N": 1, "O": 1, "P": 3, "Q": 10, "R": 1, "S": 1, "T": 1, "U": 1, "V": 4, "W": 4, "X": 8, "Y": 4, "Z": 10}

// Score returns scrabble score for given input string
func Score(input string) int {
	score := 0

	if len(input) == 0 {
		return score
	}

	splitInput := strings.Split(strings.ToUpper(input), "")

	for _, inputLetter := range splitInput {
		for mapLetter, letterValue := range scoreMap {
			if mapLetter == inputLetter {
				score += letterValue
			}
		}
	}

	return score
}
