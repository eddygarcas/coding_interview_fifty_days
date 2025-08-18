package main

// This program determines if a binary tree has a root-to-leaf path that sums to a target value.
// It uses recursive depth-first traversal, subtracting node values from the target as it goes down the tree.

import (
	"fmt"
	"time"
)

// TreeNode represents a node in a binary tree with an integer value and left/right child pointers
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// main creates a sample binary tree and tests if it contains a path summing to a target value
func main() {
	// Create binary tree:
	//        2
	//      /   \
	//     4     7
	//          / \
	//         12  15
	//        /    /
	//       20   18
	//        \
	//        28
	root := TreeNode{val: 2}
	root.left = &TreeNode{val: 4}
	root.right = &TreeNode{val: 7}
	root.right.left = &TreeNode{val: 12}
	root.right.left.left = &TreeNode{val: 20}
	root.right.left.left.right = &TreeNode{val: 28}
	root.right.right = &TreeNode{val: 15}
	root.right.right.left = &TreeNode{val: 18}

	target := 69
	start := time.Now()
	result := sumPath(&root, target)
	fmt.Printf("The tree has the sum of %v?: %v\n", target, result)
	elapsed := time.Since(start)
	fmt.Printf("The call took %s\n", elapsed)
}

// sumPath recursively checks if there exists a root-to-leaf path summing to target
// Returns true if such a path exists, false otherwise
func sumPath(root *TreeNode, target int) bool {
	if root == nil {
		return target == 0
	}
	return sumPath(root.left, target-root.val) || sumPath(root.right, target-root.val)
}
