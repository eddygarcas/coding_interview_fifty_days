package main

import (
	"fmt"
	"time"
)

//In the context of binary trees, DFS (Depth-First Search) and BFS (Breadth-First Search)
//are two fundamental tree traversal algorithms — each explores the tree in a different order:
//DFS explores as deep as possible down one branch before backtracking.
//BFS visits nodes level by level, from top to bottom and left to right.

//| Feature     | DFS                           | BFS                        |
//| ----------- | ----------------------------- | -------------------------- |
//| Approach    | Go deep first                 | Visit level by level       |
//| Data Struct | Stack / Recursion             | Queue                      |
//| Memory      | Less if tree is deep          | More if tree is wide       |
//| Use Cases   | Path finding, tree properties | Shortest path, level order |

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	//        3
	//      /   \
	//     9     20
	//          /  \
	//         15   25
	root := &TreeNode{val: 3}
	root.left = &TreeNode{val: 9}
	root.right = &TreeNode{val: 20}
	root.right.left = &TreeNode{val: 15}
	root.right.right = &TreeNode{val: 25}

	start := time.Now()
	res := levelOrder(root)
	fmt.Printf("TreeOrder level: %v\n", res)
	elapsed := time.Since(start)
	fmt.Printf("elapsed %s\n", elapsed)
}

// levelOrder performs a Breadth-First Search (BFS) traversal of the binary tree
// It uses a queue to visit nodes level by level, from left to right:
// 1. Start with root node in queue
// 2. For each level:
//   - Get current level size
//   - Process all nodes in current level
//   - Add their children to queue for next level
//
// 3. Return array of arrays where each inner array represents one level
// Time: O(n) where n is number of nodes
// Space: O(w) where w is maximum width of tree
//        3     Level 0: [3]
//      /   \
//     9     20  Level 1: [9,20]
//          /  \
//         15   25 Level 2: [15,25]

// After BFS traversal, result will be: [[3], [9,20], [15,25]]
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := [][]int{}

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.val)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		result = append(result, level)
	}
	return result
}
