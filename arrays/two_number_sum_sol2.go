package arrays

// Category: arrays
// Difficulty: Easy

// Given a non-empty array of disting integers and a resulting sum, return the two numbers that sum up to the result
// if no such numbers exist, return an empty array

func TwoNumberSumSol2(array []int, target int) []int {

	resultNumbers := []int{}
	storeMap := make(map[int]bool)

	for i := 0; i < len(array); i++ {
		result := target - array[i]
		if storeMap[result] == true {
			resultNumbers = append(resultNumbers, array[i], result)
			break
		} else {
			storeMap[array[i]] = true
		}

	}
	return resultNumbers
}
