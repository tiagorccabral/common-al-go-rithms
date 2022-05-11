package greedy

// Category: graphs
// Difficulty: Easy

// queries = [3,2,1,2,6]
// output: 17

import (
	"sort"
)

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	queriesSize := len(queries) - 1
	minTime := 0

	for i, query := range queries {
		// every query after current query must wait current query execution time
		minTime = minTime + (queriesSize-i)*query
	}

	return minTime
}
