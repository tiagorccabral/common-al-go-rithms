package arrays

// Category: arrays
// Difficulty: Easy

// Given two arrays, determine if the second is a valid subsequence of the first.
// Ex: [1, 2, 3, 4] and [2, 4] -> should return true because [2,4] is a valid subsequence of [1,2,3,4]

func IsValidSubsequence(array []int, sequence []int) bool {
	sequenceIdx := 0
	arrayIdx := 0
	isSubsequence := false

	for arrayIdx < len(array) && sequenceIdx < len(sequence) {
		if array[arrayIdx] == sequence[sequenceIdx] {
			sequenceIdx++
		}
		arrayIdx++
	}

	if sequenceIdx == len(sequence) {
		isSubsequence = true
	}

	return isSubsequence
}
