package graphs

// Category: graphs
// Difficulty: Medium

// Given a matrix of 1s and 0s, where 1s adjecent to each other represent a continuous river (diagonal do not count as a continuous river).
// Should return an array with all the lengths of possible rivers.

// Example:
// Input: matrix = [
// 	[1, 0, 0, 1, 0],
// 	[1, 0, 1, 0, 0],
// 	[0, 0, 1, 0, 1],
// 	[1, 0, 1, 0, 1],
// 	[1, 0, 1, 1, 0]
// ]

// Output: [1, 2, 2, 2, 5]

func RiverSizes(matrix [][]int) []int {
	sizes := []int{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			}
			sizes = traverseRivers(i, j, matrix, visited, sizes)
		}
	}
	return sizes
}

func traverseRivers(i int, j int, matrix [][]int, visited [][]bool, sizes []int) []int {

	riversQueue := [][]int{{i, j}}
	currentRiverSize := 0

	for len(riversQueue) > 0 {
		currentNode := riversQueue[0]
		riversQueue = riversQueue[1:]

		currentNodeI := currentNode[0]
		currentNodeJ := currentNode[1]

		if visited[currentNodeI][currentNodeJ] {
			continue
		}
		visited[currentNodeI][currentNodeJ] = true
		if matrix[currentNodeI][currentNodeJ] == 0 {
			continue
		}
		currentRiverSize += 1

		unvisitedNeighbors := getUnvisitedNeighbors(currentNodeI, currentNodeJ, visited, matrix)
		for _, neighbor := range unvisitedNeighbors {
			riversQueue = append(riversQueue, neighbor)
		}
	}

	if currentRiverSize > 0 {
		sizes = append(sizes, currentRiverSize)
	}

	return sizes
}

func getUnvisitedNeighbors(i int, j int, visited [][]bool, matrix [][]int) [][]int {
	unvisited := [][]int{}

	if j > 0 && !visited[i][j-1] {
		unvisited = append(unvisited, []int{i, j - 1})
	}
	if j < (len(matrix[0])-1) && !visited[i][j+1] {
		unvisited = append(unvisited, []int{i, j + 1})
	}
	if i > 0 && !visited[i-1][j] {
		unvisited = append(unvisited, []int{i - 1, j})
	}
	if i < len(matrix)-1 && !visited[i+1][j] {
		unvisited = append(unvisited, []int{i + 1, j})
	}

	return unvisited
}
