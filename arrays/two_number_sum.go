package arrays

// Category: arrays
// Difficulty: Easy

// Given a non-empty array of disting integers and a resulting sum, return the two numbers that sum up to the result
// if no such numbers exist, return an empty array

import "sort"

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	leftIndex := 0
	rightIndex := len(array) - 1
	result := []int{}

	for i := 0; i < len(array); i++ {
		tempSum := array[leftIndex] + array[rightIndex]

		if tempSum == target {
			result = append(result, array[leftIndex], array[rightIndex])
			break
		} else if tempSum < target {
			leftIndex = leftIndex + 1
		} else if tempSum > target {
			rightIndex = rightIndex - 1
		}
	}

	return result
}
