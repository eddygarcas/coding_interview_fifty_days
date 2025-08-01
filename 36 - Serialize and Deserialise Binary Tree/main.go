package main

// This program implements binary tree serialization and deserialization.
// It converts a binary tree to a string format and back, using a preorder traversal approach.
// 'X' is used to denote null nodes, and '#' is used as a delimiter between nodes.

import (
	"fmt"
	"strconv"
	"time"
)

// TreeNode represents a node in a binary tree
// It contains an integer value, pointers to left and right children,
// and a string field to store the serialized representation
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
	ans   string
}

func main() {
	// Create sample binary tree:
	//        1
	//      /   \
	//     2     3
	//          / \
	//         4   5
	root := TreeNode{val: 1}
	root.left = &TreeNode{val: 2}
	root.right = &TreeNode{val: 3}
	root.right.left = &TreeNode{val: 4}
	root.right.right = &TreeNode{val: 5}

	serialized := []byte("12XX34XX5XX")

	start := time.Now()
	root.serialize(&root)
	fmt.Printf("Serialized tree %v\n", root.ans)
	//var ans *TreeNode
	//ans := &TreeNode{}
	ans, _ := root.deserialize(serialized)
	ans.printTree()
	elapsed := time.Since(start)
	fmt.Printf("\nThe call took %s\n", elapsed)
}

// serialize converts the binary tree to a string representation
// using preorder traversal (root-left-right)
// Each node's value is followed by '#' and null nodes are represented as 'X#'
func (t *TreeNode) serialize(root *TreeNode) {
	if root == nil {
		t.ans += "X#"
		return
	}
	t.ans += strconv.Itoa(root.val) + "#"
	t.serialize(root.left)
	t.serialize(root.right)
	return
}

// deserialize reconstructs a binary tree from its serialized byte representation
// Returns the reconstructed tree node and remaining serialized data
func (t *TreeNode) deserialize(serialized []byte) (*TreeNode, []byte) {
	if len(serialized) == 0 {
		return nil, serialized
	}

	val := serialized[0]
	serialized = serialized[1:]
	fmt.Printf("value %c and serialized %s\n", val, serialized)

	if val == 'X' {
		return nil, serialized
	}
	node := &TreeNode{}
	node.val, _ = strconv.Atoi(string(val))
	node.printTree()
	fmt.Printf("\n")

	node.left, serialized = t.deserialize(serialized)
	node.right, serialized = t.deserialize(serialized)
	return node, serialized
}

// printTree performs a preorder traversal of the tree and prints each node's value
// Null nodes are printed as 'X'
func (t *TreeNode) printTree() {
	if t == nil {
		fmt.Print("X ")
		return
	}
	fmt.Printf("%d ", t.val)
	t.left.printTree()
	t.right.printTree()
}
