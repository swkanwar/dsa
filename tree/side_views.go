package main

// LeftSideView returns the values of the nodes you can see standing on the left side of the tree.
func LeftSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			// If it is the first node of the current level, add it to the result.
			if i == 0 {
				result = append(result, curr.Val)
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}

	return result
}

// RightSideView returns the values of the nodes you can see standing on the right side of the tree.
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			// If it is the last node of the current level, add it to the result.
			if i == levelSize-1 {
				result = append(result, curr.Val)
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}

	return result
}

// TopView returns the values of the nodes visible from the top of the tree.
func TopView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	type element struct {
		node *TreeNode
		hd   int // horizontal distance
	}

	m := make(map[int]int)
	queue := []element{{root, 0}}
	minHd, maxHd := 0, 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := m[curr.hd]; !ok {
			m[curr.hd] = curr.node.Val
		}

		if curr.hd < minHd {
			minHd = curr.hd
		}
		if curr.hd > maxHd {
			maxHd = curr.hd
		}

		if curr.node.Left != nil {
			queue = append(queue, element{curr.node.Left, curr.hd - 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, element{curr.node.Right, curr.hd + 1})
		}
	}

	var result []int
	for i := minHd; i <= maxHd; i++ {
		result = append(result, m[i])
	}
	return result
}

// BottomView returns the values of the nodes visible from the bottom of the tree.
func BottomView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	type element struct {
		node *TreeNode
		hd   int
	}

	m := make(map[int]int)
	queue := []element{{root, 0}}
	minHd, maxHd := 0, 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Update map every time so the last node at each hd (deepest) remains
		m[curr.hd] = curr.node.Val

		if curr.hd < minHd {
			minHd = curr.hd
		}
		if curr.hd > maxHd {
			maxHd = curr.hd
		}

		if curr.node.Left != nil {
			queue = append(queue, element{curr.node.Left, curr.hd - 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, element{curr.node.Right, curr.hd + 1})
		}
	}

	var result []int
	for i := minHd; i <= maxHd; i++ {
		result = append(result, m[i])
	}
	return result
}
