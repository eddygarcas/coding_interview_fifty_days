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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := [][]int{}

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
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
