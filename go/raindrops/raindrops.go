package raindrops

import (
	"strconv"
)

//Convert converts number to string if it has any of the given numbers as a factor
func Convert(input int) string {

	var raindropSpeak string

	if input%3 == 0 {
		raindropSpeak += "Pling"
	}

	if input%5 == 0 {
		raindropSpeak += "Plang"
	}

	if input%7 == 0 {
		raindropSpeak += "Plong"
	}

	if len(raindropSpeak) == 0 {
		raindropSpeak = strconv.Itoa(input)
	}

	return raindropSpeak
}
