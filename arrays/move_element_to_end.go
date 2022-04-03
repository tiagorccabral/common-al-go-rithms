package arrays

// Category: arrays
// Difficulty: Medium

// Given an array of integers and an integer, move all values equal to the integer to the end of the array (in place)

// Example Input:
// {
//   "array": [2, 4, 2, 5, 6, 2, 2],
//   "toMove": 2
// }

// Expected Output: [6, 4, 5, 2, 2, 2, 2]

// Solution complexity: O(n)

func MoveElementToEnd(array []int, toMove int) []int {

	initialPointer := 0
	endPointer := len(array) - 1

	for initialPointer < endPointer {
		if array[initialPointer] == toMove && array[endPointer] != toMove {
			tmp := array[initialPointer]
			array[initialPointer] = array[endPointer]
			array[endPointer] = tmp
		}
		if array[initialPointer] != toMove {
			initialPointer++
		}
		if array[endPointer] == toMove {
			endPointer--
		}
	}

	return array
}
