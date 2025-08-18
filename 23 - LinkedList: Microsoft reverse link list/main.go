package main

import (
	"fmt"
	"time"
)

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int       // Value stored in the node
	Next *ListNode // Pointer to the next node
}

// reverse attempts to reverse a linked list from node f to node l (non-standard implementation).
// It prints the values at the start and end of the reversal.
// This function is not a conventional in-place linked list reversal and is likely for demonstration.
func reverse(f, l *ListNode) {
	// Store the original first node to reset the loop
	origin := f
	for f != nil {
		// Check if we've reached the last node in the reversal
		if f.Next == l {
			// Print the values at the start and end of the reversal
			fmt.Printf("First: %v  Last: %v\n", f.Val, f.Next.Val)
			// Reverse the link between the last two nodes
			l.Next = f
			// Set the next pointer of the new last node to nil
			f.Next = nil
			// Update the last node pointer
			l = f
			// Reset the loop to the original first node
			f = origin
		} else {
			// Move to the next node in the list
			f = f.Next
		}
	}
}

// reverseList reverses a singly linked list in-place.
// Takes the head of the list and returns the new head after reversal.
// Prints the value of each node as it is processed.
func reverseList(head *ListNode) *ListNode {
	// Initialize the previous node pointer to nil
	var prev *ListNode = nil
	// Initialize the current node pointer to the head of the list
	current := head
	for current != nil {
		// Store the next node in the list before updating the current node's next pointer
		nextNode := current.Next
		// Print the value of the current node
		fmt.Printf("First: %v  Last: %v\n", current.Val, current.Val)
		// Reverse the link between the current node and the previous node
		current.Next = prev
		// Update the previous node pointer
		prev = current
		// Move to the next node in the list
		current = nextNode
	}
	// Return the new head of the reversed list
	return prev
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}

	// Create a cycle: node6.Next points back to node3
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = nil

	f := node1
	l := node6
	start := time.Now()
	reverse(f, l)
	for l != nil {
		fmt.Printf("%v,", l.Val)
		l = l.Next
	}
	elapsed := time.Since(start)
	fmt.Printf("\nelapsed: %s\n", elapsed)

	start = time.Now()
	f = node6
	rev := reverseList(f)
	for rev != nil {
		fmt.Printf("%v,", rev.Val)
		rev = rev.Next
	}
	elapsed = time.Since(start)
	fmt.Printf("\nelapsed: %s\n", elapsed)
}
