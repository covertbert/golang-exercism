// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	s "strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var response string

	switch {
	case endsWith(remark, "?"):
		switch {
		case isAllUpperCase(remark):
			response = "Calm down, I know what I'm doing!"
		default:
			response = "Sure."
		}
	case isAllUpperCase(remark):
		response = "Whoa, chill out!"
	case isBlank(remark):
		response = "Fine. Be that way!"
	default:
		response = "Whatever."
	}

	return response
}

func endsWith(remark string, character string) bool {
	trimmedRemark := s.TrimSpace(remark)

	if len(trimmedRemark) <= 1 {
		return false
	}

	return trimmedRemark[len(trimmedRemark)-1:] == character
}

func isAllUpperCase(remark string) bool {
	if s.ToUpper(remark) == s.ToLower(remark) {
		return false
	}
	return s.ToUpper(remark) == remark
}

func isBlank(remark string) bool {
	trimmedRemark := s.TrimSpace(remark)
	return len(trimmedRemark) == 0
}
