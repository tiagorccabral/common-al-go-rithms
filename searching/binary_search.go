package searching

// Category: Searching
// Difficulty: Easy

// Given an ordered array of numbers and a target value, return the index of the value in the ordered array
// If the array does not have the target number, return -1

func BinarySearch(array []int, target int) int {

	start := 0
	end := len(array) - 1

	resultIdx := -1

	for start <= end {
		middle := (start + end) / 2
		if target == array[middle] {
			resultIdx = middle
			break
		} else if target < array[middle] {
			end = middle - 1
		} else if target > array[middle] {
			start = middle + 1
		}
	}

	return resultIdx
}
