package main

import (
	"fmt"
)

// TreeNode defines a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// VerticalTraversal returns the vertical order traversal of a binary tree.
func VerticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// columnMap stores values grouped by their horizontal distance from root.
	columnMap := make(map[int][]int)

	// Queue for BFS: stores node and its column index.
	type queueItem struct {
		node *TreeNode
		col  int
	}
	queue := []queueItem{{root, 0}}

	minCol, maxCol := 0, 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		node, col := curr.node, curr.col
		columnMap[col] = append(columnMap[col], node.Val)

		if col < minCol {
			minCol = col
		}
		if col > maxCol {
			maxCol = col
		}

		if node.Left != nil {
			queue = append(queue, queueItem{node.Left, col - 1})
		}
		if node.Right != nil {
			queue = append(queue, queueItem{node.Right, col + 1})
		}
	}

	// Prepare the result by iterating from the leftmost column to the rightmost.
	result := make([][]int, 0, len(columnMap))
	for i := minCol; i <= maxCol; i++ {
		if val, ok := columnMap[i]; ok {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	// Example Tree:
	//      3
	//     / \
	//    9   20
	//       /  \
	//      15   7
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	result := VerticalTraversal(root)
	fmt.Printf("Vertical Traversal: %v\n", result) // Expected: [[9] [3 15] [20] [7]]
}
