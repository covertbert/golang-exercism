package reverse

// String reverses a string
func String(input string) string {
	chars := []rune(input)

	for index, letter := 0, len(chars)-1; index < letter; index, letter = index+1, letter-1 {
		chars[index], chars[letter] = chars[letter], chars[index]
	}

	return string(chars)
}
