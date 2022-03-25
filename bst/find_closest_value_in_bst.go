package bst

import "math"

// Category: bst
// Difficulty: Easy

// Given a BST and a target value, find the closest value to target in the BST.

// Example Tree =   10
//			 					/		\
//			 				5				15
//			 			/				/		 \
//			 		2				13			22
//			 	/						\
//			 1						  14

// target = 12

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

var ClosestValue int = math.MaxInt64

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (tree *BST) FindClosestValue(target int) int {

	currentMinDiff := Abs(Abs(target) - Abs(ClosestValue))
	currentDiff := Abs(Abs(target) - Abs(tree.Value))

	if currentDiff < currentMinDiff {
		ClosestValue = tree.Value
	}

	if currentDiff != 0 {
		if tree.Value < target && tree.Right != nil {
			tree.Right.FindClosestValue(target)
		} else if tree.Value > target && tree.Left != nil {
			tree.Left.FindClosestValue(target)
		}
	}

	return ClosestValue
}
