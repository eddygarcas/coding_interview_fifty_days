package main

// This program finds the lowest common ancestor (LCA) of two nodes in a binary tree.
// The LCA is the lowest node in the tree that has both nodes as descendants.
// It uses a recursive depth-first search approach.

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

	start := time.Now()
	lca := lcanode(&root, 28, 4)
	fmt.Printf("lca = %v\n", lca.val)
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
}

// lcanode finds the lowest common ancestor of two nodes with values p and q in the binary tree
// Returns the LCA node if found, nil otherwise
// Uses post-order traversal where:
// - If current node matches either value, return it
// - Recursively search left and right subtrees
// - If both subtrees return non-nil, current node is LCA
// - If one subtree returns nil, return the non-nil subtree result
func lcanode(root *TreeNode, p, q int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.val == p || root.val == q {
		fmt.Printf("value found: %d\n", root.val)
		return root
	}
	left := lcanode(root.left, p, q)
	right := lcanode(root.right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
