package main

import "fmt"

// ListNode defines a node in a singly linked list.
type ListNode struct {
	Val  int       // Value stored in the node
	Next *ListNode // Pointer to the next node in the list
}

func main() {
	// Create nodes for the linked list: 1 -> 2 -> 3 -> 4 -> 5 -> 6
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}

	// Link the nodes together
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = nil // End of list

	// Remove the 2nd node from the end (should remove node 5)
	result := removeNthNode(node1, 2)
	for result != nil {
		fmt.Printf("%d,", result.Val)
		result = result.Next
	}
}

// removeNthNode removes the nth node from the end of the singly linked list.
// t: Head of the linked list
// n: The position from the end (1-based index)
// Returns the head of the modified list.
func removeNthNode(t *ListNode, n int) *ListNode {
	temp := t  // Pointer to traverse the list
	index := 0 // Total number of nodes in the list
	ans := t
	// First pass: count the total number of nodes
	for temp != nil {
		index++
		temp = temp.Next
	}
	// Second pass: move to the node just before the one to remove
	// Rather than return the ones to delete, return the previous ones.
	length := index - n
	for length > 0 {
		ans = ans.Next
		length--
	}
	ans.Next = ans.Next.Next
	return t
}
