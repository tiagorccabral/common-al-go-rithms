package arrays

import (
	"math"
	"sort"
)

// Category: arrays
// Difficulty: Medium

// Given two arrays of integers, find and return the pair of numbers (one from each array) that has the smallest absolute difference

// Example:
// [-1, 5, 10, 20, 28, 3]
// [26, 134, 135, 15, 17]
// Output: [28, 26]

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SmallestDifference(array1, array2 []int) []int {
	SmallestAbsDiff := math.MaxInt64
	resultNumbers := []int{}

	sort.Ints(array1)
	sort.Ints(array2)

	array1Idx := 0
	array2Idx := 0

	for array1Idx < len(array1) && array2Idx < len(array2) {
		difference := Abs(array1[array1Idx] - array2[array2Idx])
		if difference == 0 {
			resultNumbers = []int{array1[array1Idx], array2[array2Idx]}
			break
		} else if difference < SmallestAbsDiff {
			SmallestAbsDiff = difference
			resultNumbers = []int{array1[array1Idx], array2[array2Idx]}
		}

		if array1[array1Idx] < array2[array2Idx] {
			array1Idx++
		} else if array1[array1Idx] > array2[array2Idx] {
			array2Idx++
		}
	}

	return resultNumbers
}
