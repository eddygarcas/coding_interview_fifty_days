package main

import (
	"fmt"
	"time"
)

// This program implements post-order traversal of a binary tree.
// In post-order traversal, for each node we:
// 1. Traverse left subtree
// 2. Traverse right subtree
// 3. Visit the node
//
// Example tree:
//        3
//      /   \
//     9    20
//         /  \
//        15   7
//
// Post-order traversal output: [9,15,7,20,3]
// (left subtree -> right subtree -> root)

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
	res := levelPostOrder(root)
	fmt.Printf("TreeOrder level: %v\n", res)
	elapsed := time.Since(start)
	fmt.Printf("elapsed %s\n", elapsed)
}

// levelPostOrder performs iterative post-order traversal using two stacks
// Stack states visualization:
// Initial:
// stack1: [3]
// stack2: []

// After first pop from stack1:
// stack1: [9, 20]  // Left and right children of 3
// stack2: [3]

// After second pop from stack1:
// stack1: [20]
// stack2: [3, 9]

// After third pop from stack1:
// stack1: [15, 7]  // Children of 20
// stack2: [3, 9, 20]

// After fourth pop from stack1:
// stack1: [7]
// stack2: [3, 9, 20, 15]

// After fifth pop from stack1:
// stack1: []
// stack2: [3, 9, 20, 15, 7]

// Final result after processing stack2:
// result: [9, 15, 7, 20, 3]
// Time: O(n) where n is number of nodes
// Space: O(n) for the two stacks
func levelPostOrder(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	var stack1 []*TreeNode
	var stack2 []*TreeNode
	stack1 = append(stack1, root)

	for len(stack1) > 0 {
		x := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, x)
		if x.left != nil {
			stack1 = append(stack1, x.left)
		}
		if x.right != nil {
			stack1 = append(stack1, x.right)
		}
	}
	for len(stack2) > 0 {
		y := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		result = append(result, y.val)

	}
	return result
}
