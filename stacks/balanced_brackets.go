package stacks

// Category: stacks
// Difficulty: Medium

// Given a string containing a number of bracks, determine if the brackets are balanced or not.

// Input: "([])(){}(())()()"
// output: true

var openingChars = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
}

var closingChars = map[rune]bool{
	')': true,
	']': true,
	'}': true,
}

var equivalentMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func BalancedBrackets(s string) bool {
	stack := []rune{}

	for _, currentChar := range s {

		if _, found := closingChars[currentChar]; found {
			if len(stack) > 0 && currentChar == equivalentMap[stack[len(stack)-1]] {
				stack = stack[0 : len(stack)-1]
				continue
			} else {
				return false
			}
		} else if _, found := openingChars[currentChar]; found {
			stack = append(stack, currentChar)
		}
	}

	return len(stack) == 0
}
