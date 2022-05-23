package strings

// Category: strings
// Difficulty: Easy

// Given a string, return the index of the first non repeating string

// ex: "abcdcaf"
// output: 1

func FirstNonRepeatingCharacter(str string) int {
	occCount := map[rune]int{}

	for _, val := range str {
		occCount[val] = occCount[val] + 1
	}

	for idx, val := range str {
		if count, ok := occCount[val]; ok && count == 1 {
			return idx
		}
	}

	return -1
}
