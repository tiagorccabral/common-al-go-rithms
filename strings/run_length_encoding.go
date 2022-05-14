package strings

// Category: strings
// Difficulty: Easy

// Given a string of chars from A-Z, return the run-length encoding of that string.
// Obs, if there's more than 9 consecutive chars, they should be split in to two or more groups of 9
// Ex input: AAAAAAAAAAAAABBCCCCDD
// Output:   9A4A2B4C2D

import (
	"strconv"
)

func RunLengthEncoding(str string) string {
	currentChar := string(str[0])
	currentCharCount := 1
	finalString := ""
	for i := 1; i < len(str); i++ {
		if currentChar != string(str[i]) || currentCharCount == 9 {

			finalString += strconv.Itoa(currentCharCount)
			finalString += string(currentChar)

			currentChar = string(str[i])
			currentCharCount = 1
		} else {
			currentCharCount += 1
		}
	}

	finalString += strconv.Itoa(currentCharCount)
	finalString += string(currentChar)
	return finalString
}
