package main

import (
	"fmt"
	"time"
)

// TreeNode represents a node in a binary tree with a value and left/right child pointers
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// main creates a sample symmetric binary tree and tests if it is symmetric
func main() {
	// Create a symmetric tree:
	//      1
	//    /   \
	//   2     2
	//  / \   / \
	// 3   4 4   3
	root := TreeNode{val: 1}
	root.left = &TreeNode{val: 2}
	root.right = &TreeNode{val: 2}
	root.left.left = &TreeNode{val: 3}
	root.left.right = &TreeNode{val: 4}
	root.right.left = &TreeNode{val: 4}
	root.right.right = &TreeNode{val: 3}
	start := time.Now()
	fmt.Println(isSymetric(&root, &root))
	elapsed := time.Since(start)
	fmt.Printf("The call took %s\n", elapsed)
}

// isSymetric checks if two binary trees are mirror images of each other
// by recursively comparing their left and right subtrees
func isSymetric(t1 *TreeNode, t2 *TreeNode) bool {
	// Base cases: both nodes are nil or one node is nil
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	// Recursive check for mirror symmetry
	mirrorLeft := isSymetric(t1.right, t2.left)
	mirrorRight := isSymetric(t1.left, t2.right)
	return mirrorLeft && mirrorRight && t1.val == t2.val
}
