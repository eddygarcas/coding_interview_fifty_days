package main

// This program finds the maximum path sum in a binary tree.
// A path is a sequence of nodes where each pair of adjacent nodes has a parent-child connection.
// The path sum is the sum of the node values along the path.
// The path may start and end at any node and doesn't need to go through the root.

import (
	"fmt"
	"math"
	"time"
)

// TreeNode represents a node in a binary tree.
// It contains an integer value, pointers to left and right children,
// and a field to store the maximum path sum found so far.
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
	ans   int
}

func main() {
	// Create binary tree:
	//       -10
	//      /    \
	//     9     20
	//          /  \
	//         15   7
	// Expected max path sum: 42 (20 + 15 + 7)
	root := TreeNode{val: -10, ans: math.MinInt32}
	root.left = &TreeNode{val: 9}
	root.right = &TreeNode{val: 20}
	root.right.left = &TreeNode{val: 15}
	root.right.right = &TreeNode{val: 7}

	start := time.Now()
	ans := &TreeNode{}
	ans.getMaxPathSum(&root)
	fmt.Printf("Max path sum %d\n", ans.ans)
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed %s\n", elapsed)
}

// getMaxPathSum calculates the maximum path sum in the binary tree recursively.
// For each node, it computes:
// 1. Maximum path sum including the node and at most one child (returned up)
// 2. Maximum path sum including the node and both children (updates global max)
// Returns the maximum path sum that can be extended by parent nodes
func (t *TreeNode) getMaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := t.getMaxPathSum(root.left)
	rightSum := t.getMaxPathSum(root.right)

	maxSide := max(root.val, root.val+max(leftSum, rightSum))
	maxCurrent := max(maxSide, root.val+leftSum+rightSum)
	t.ans = max(t.ans, maxCurrent)
	return maxSide
}
