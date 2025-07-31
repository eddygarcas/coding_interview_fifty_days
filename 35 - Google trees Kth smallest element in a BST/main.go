package main

// This program finds the kth smallest element in a Binary Search Tree (BST).
// It uses an in-order traversal approach since BST in-order traversal gives elements in ascending order.
// The algorithm maintains a counter that decrements with each node visited until reaching the kth element.
//
// For example, given this BST:
//                          130
//                /                   \
//              70                    190
//           /      \              /       \
//         40       100          160        220
//       /    \    /    \      /    \     /    \
//     20     55  80    110  140    170  200   240
//    /  \        \      \     \      \    \   /  \
//   5   15       90     120   150    180  210 230 250
//
// And k=7, the program will return 90 (the 7th smallest element)
//
// Time Complexity: O(n) where n is number of nodes
// Space Complexity: O(h) where h is height of tree due to recursion stack

import (
	"fmt"
	"time"
)

// TreeNode represents a node in a binary search tree
type TreeNode struct {
	val   int       // Value stored in the node
	left  *TreeNode // Pointer to left child node
	right *TreeNode // Pointer to right child node
	ans   int       // Stores the kth smallest element when found
}

// main creates a sample BST and finds its kth smallest element
func main() {
	// Construct the BST
	root := TreeNode{val: 130}

	root.left = &TreeNode{val: 70}
	root.left.left = &TreeNode{val: 40}
	root.left.left.left = &TreeNode{val: 20}
	root.left.left.left.left = &TreeNode{val: 5}
	root.left.left.left.right = &TreeNode{val: 15}
	root.left.left.right = &TreeNode{val: 55}
	root.left.left.right.right = &TreeNode{val: 65}
	root.left.right = &TreeNode{val: 100}
	root.left.right.left = &TreeNode{val: 80}
	root.left.right.left.right = &TreeNode{val: 90}
	root.left.right.right = &TreeNode{val: 110}
	root.left.right.right.right = &TreeNode{val: 120}

	root.right = &TreeNode{val: 190}
	root.right.left = &TreeNode{val: 160}
	root.right.left.left = &TreeNode{val: 140}
	root.right.left.left.right = &TreeNode{val: 150}
	root.right.left.right = &TreeNode{val: 170}
	root.right.left.right.right = &TreeNode{val: 180}
	root.right.right = &TreeNode{val: 220}
	root.right.right.left = &TreeNode{val: 200}
	root.right.right.left.right = &TreeNode{val: 210}
	root.right.right.right = &TreeNode{val: 240}
	root.right.right.right.left = &TreeNode{val: 230}
	root.right.right.right.right = &TreeNode{val: 250}

	start := time.Now()
	k := 7
	t := TreeNode{}
	t.smallestElement(&root, &k)
	fmt.Printf("Sorted array element is %v\n", t.ans)
	elapsed := time.Since(start)
	fmt.Printf("The call took %s and complexity Time/Space O(n)\n", elapsed)
}

// smallestElement performs an in-order traversal to find the kth smallest element
// Parameters:
//   - root: pointer to current node in traversal
//   - k: pointer to remaining count until kth element is found
//
// The method stores result in the receiver's ans field when k reaches 0
func (t *TreeNode) smallestElement(root *TreeNode, k *int) {
	if root == nil {
		return
	}
	// Traverse left subtree
	t.smallestElement(root.left, k)
	// Process current node
	*k--
	if *k == 0 {
		t.ans = root.val
		return
	}
	// Traverse right subtree
	t.smallestElement(root.right, k)
}
