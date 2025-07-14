package main

import "fmt"

// ListNode defines a node in a singly linked list.
// Each node contains an integer value and a pointer to the next node in the list.
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

	// Link the nodes together to form a singly linked list
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = nil // End of list

	// Remove the 2nd node from the end (should remove node 5)
	// Pass the head of the list and the position from the end (1-based index)
	result := node1
	node1.removeNthNode(2)
	// Print the modified list
	for result != nil {
		fmt.Printf("%d,", result.Val)
		result = result.Next
	}
}

// removeNthNode removes the nth node from the end of the singly linked list.
// Parameters:
// t: Head of the linked list
// n: The position from the end (1-based index)
// Returns the head of the modified list.
func (t *ListNode) removeNthNode(n int) {
	// temp is used to count the total number of nodes in the list
	temp := t
	index := 0 // Total number of nodes in the list
	ans := t   // Used to traverse to the node before the one to remove

	// First pass: count the total number of nodes in the list
	for temp != nil {
		index++
		temp = temp.Next
	}

	// Second pass: move to the node just before the one to remove
	// We subtract n+1 from the total count to get the position from the start
	for l := 0; l < index-(n+1); l++ {
		ans = ans.Next
	}
	// Remove the nth node from the end by adjusting pointers
	// We set the next pointer of the node before the one to remove to nil
	ans.Next = nil
}
