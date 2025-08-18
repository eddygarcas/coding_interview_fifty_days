package main

import (
	"fmt"
	"time"
)

// TreeNode represents a node in a binary tree with a value and left/right child pointers
type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

// main creates a sample binary tree and calculates its maximum depth
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
	//        21
	root := TreeNode{Val: 2}
	root.left = &TreeNode{Val: 4}
	root.right = &TreeNode{Val: 7}
	root.right.left = &TreeNode{Val: 12}
	root.right.left.left = &TreeNode{Val: 20}
	root.right.left.left.right = &TreeNode{Val: 21}
	root.right.right = &TreeNode{Val: 15}
	root.right.right.left = &TreeNode{Val: 18}
	start := time.Now()
	res := depthOfBinaryTree(&root)
	fmt.Printf("Max depth of binary tree in %d\n", res)
	elapsed := time.Since(start)
	fmt.Printf("The call took %s\n", elapsed)
}

// depthOfBinaryTree calculates the maximum depth of a binary tree recursively
// Returns 0 for empty tree, otherwise returns max depth of left or right subtree + 1
func depthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + depthOfBinaryTree(root.left)
	right := 1 + depthOfBinaryTree(root.right)
	fmt.Printf("Node %d has left %d and right %d\n", root.Val, left, right)

	return max(left, right)
}
