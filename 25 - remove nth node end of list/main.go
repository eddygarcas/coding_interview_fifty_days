package main

import (
	"fmt"
	"time"
)

// ListNode defines a node in a singly linked list.
// Each node contains an integer value and a pointer to the next node.
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

	// Link the nodes together to form the list
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = nil // End of list

	// Remove the 2nd node from the end (should remove node 5)
	// Calls the method on the head node
	start := time.Now()
	result := node1
	node1.removeNthNode(2)
	elapsed := time.Since(start)
	// Print the modified list
	for result != nil {
		fmt.Printf("%d,", result.Val)
		result = result.Next
	}
	fmt.Printf("\nTime elapsed: %s\n", elapsed)
}

// removeNthNode removes the nth node from the end of the singly linked list.
// This is a method on the ListNode type.
// Parameters:
// n: The position from the end (1-based index)
// Modifies the list in place. No return value.
func (t *ListNode) removeNthNode(n int) {
	// temp is used to count the total number of nodes in the list
	temp := t
	index := 0 // Total number of nodes in the list
	ans := t   // Used to traverse to the node before the one to remove

	// First pass: count the total number of nodes in the list
	// Traverse the list to get the total count of nodes
	for temp != nil {
		index++
		temp = temp.Next
	}

	// Second pass: move to the node just before the one to remove
	// We subtract n+1 from the total count to get the position from the start
	// Traverse the list to get to the node before the one to remove
	for l := 0; l < index-(n+1); l++ {
		ans = ans.Next
	}
	// Remove the nth node from the end by adjusting pointers
	// We set the next pointer of the node before the one to remove to nil
	// This effectively removes the nth node from the end
	ans.Next = nil
}
