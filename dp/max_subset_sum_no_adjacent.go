package dp

// Category: dinamyc programming
// Difficulty: Medium

// Write a function that takes an array, and returns the biggest possible sum of non adjecent numbers
// Example:
// Input: [75, 105, 120, 75, 90, 135]
// Output: 333 -> 75 + 120 + 135

func maxNum(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func MaxSubsetSumNoAdjacent(array []int) int {
	maxSums := []int{}

	if len(array) == 0 {
		return 0
	}

	for i := 0; i < len(array); i++ {
		if i >= 2 {
			crtPlusLastAdjc := array[i] + maxSums[i-2]
			if maxNum(crtPlusLastAdjc, maxSums[i-1]) == crtPlusLastAdjc {
				maxSums = append(maxSums, crtPlusLastAdjc)
			} else {
				maxSums = append(maxSums, maxSums[i-1])
			}
		} else {
			if i == 0 {
				maxSums = append(maxSums, array[i])
			} else {
				if array[i] >= array[0] {
					maxSums = append(maxSums, array[i])
				} else {
					maxSums = append(maxSums, array[0])
				}
			}
		}
	}

	return maxSums[len(maxSums)-1]
}
