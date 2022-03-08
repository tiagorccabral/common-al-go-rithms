package arrays

import ("sort")

// Category: arrays
// Difficulty: Easy

// Given an array of values (coins), return the minimum amount of change that CANNOT be made with the coins

func NonConstructibleChange(coins []int) int {
	minimumChange := 0
	if (len(coins) > 0) {
		sort.Ints(coins)
		for i:=0;i<len(coins); i++ {
			if coins[i] > minimumChange + 1 {
				break
			} else {
				minimumChange = minimumChange + coins[i]
			}
		}
		// the minimum will always be what was added until the end of the loop + 1
		// in the first if we add because the next coin is greater than the current change + 1
		// in the second if we add because the minimum will be the sum of all coins + 1
		minimumChange++
		
	} else {
		minimumChange = 1
	}
	return minimumChange
}
