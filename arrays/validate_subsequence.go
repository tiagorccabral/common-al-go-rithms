package arrays

// Category: arrays
// Difficulty: Easy

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
