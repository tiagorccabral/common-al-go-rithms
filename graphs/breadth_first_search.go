package graphs

// Category: graphs
// Difficulty: Medium

// Given a graph, return an array with all the elements of the array in BFS (breadth first search) order

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	queue := []*Node{n}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		array = append(array, current.Name)

		for _, children := range current.Children {
			queue = append(queue, children)
		}
	}
	return array
}
