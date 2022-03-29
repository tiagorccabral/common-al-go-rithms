package arrays

// Category: arrays
// Difficulty: Easy

func IsValidSubsequence(array []int, sequence []int) bool {
	sequenceIdx := 0
	arrayIdx := 0

	for arrayIdx < len(array) && sequenceIdx < len(sequence) {
		if array[arrayIdx] == sequence[sequenceIdx] {
			sequenceIdx++
		}
	}
	return false
}
