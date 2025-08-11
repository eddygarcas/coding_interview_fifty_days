package main

import (
	"fmt"
	"time"
)

// This program implements zigzag level order traversal of a binary tree.
// In zigzag traversal, nodes are visited level by level, alternating between:
// - Left to right for odd numbered levels (1,3,5...)
// - Right to left for even numbered levels (2,4,6...)
//
// Example tree:
//        3        Level 1: [3] (left->right)
//      /   \
//     9    20     Level 2: [20,9] (right->left)
//         /  \
//        15   7    Level 3: [15,7] (left->right)
//
// The zigzag pattern creates output: [[3],[20,9],[15,7]]

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	root := &TreeNode{val: 3}
	root.left = &TreeNode{val: 9}
	root.right = &TreeNode{val: 20}
	root.right.left = &TreeNode{val: 15}
	root.right.right = &TreeNode{val: 7}

	start := time.Now()
	res := levelOrderZigZag(root)
	fmt.Printf("TreeOrder level: %v\n", res)
	elapsed := time.Since(start)
	fmt.Printf("elapsed %s\n", elapsed)
}

// levelOrderZigZag performs a modified BFS traversal of the binary tree
// where alternate levels are reversed to create a zigzag pattern
// Time: O(n) where n is number of nodes
// Space: O(w) where w is maximum width of tree
func levelOrderZigZag(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := [][]int{}
	// Check the level, for odd will go left->right otherwise right->left
	levelIndex := 1

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
		if levelIndex%2 == 0 {
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		result = append(result, level)
		levelIndex++
	}
	return result
}
