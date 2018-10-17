// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	s "strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var response string

	switch {
	case len(remark) == 0:
		response = "Fine. Be that way!"
	case isAllUpperCase(remark) && endsWith(remark, "?"):
		response = "Calm down, I know what I'm doing!"
	case endsWith(remark, "?"):
		response = "Sure."
	case isAllUpperCase(remark):
		response = "Whoa, chill out!"
	default:
		response = "Whatever."
	}

	return response
}

func endsWith(remark string, character string) bool {
	trimmedRemark := s.TrimSpace(remark)
	return trimmedRemark[len(trimmedRemark)-1:] == character
}

func isAllUpperCase(remark string) bool {
	if s.ToUpper(remark) == s.ToLower(remark) {
		return false
	}
	return s.ToUpper(remark) == remark
}

func isAllLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
