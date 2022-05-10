package graphs

func HasSingleCycle(array []int) bool {

	visitedCnt := 0
	currentIdx := 0

	for i := 0; i < len(array); i++ {
		if visitedCnt > 0 && currentIdx == 0 {
			return false
		}
		visitedCnt += 1

		nextIdx := (array[currentIdx] + currentIdx) % len(array)
		if nextIdx < 0 {
			nextIdx = nextIdx + len(array)
		}

		currentIdx = nextIdx
	}

	return currentIdx == 0
}
