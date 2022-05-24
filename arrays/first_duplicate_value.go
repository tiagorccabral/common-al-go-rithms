package arrays

// Category: arrays
// Difficulty: Medium

// Given an array of integers, return the value of the first duplicate value in the array
// If no duplicate is found, return -1

func FirstDuplicateValue(array []int) int {
	dupMap := map[int]int{}

	for _, val := range array {
		if _, found := dupMap[val]; found {
			return val
		}
		dupMap[val] += 1
	}

	return -1
}
