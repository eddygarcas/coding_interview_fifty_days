package main

import "fmt"

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int       // Value stored in the node
	Next *ListNode // Pointer to the next node
}

func main() {
	// Create nodes for the linked list
	node1 := &ListNode{Val: 5}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 1}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 4}

	// Link the nodes together to form the list: 5 -> 2 -> 1 -> 3 -> 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = nil

	result := node1
	node1.oddEvenList()
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}

// oddEvenList rearranges the list so that all odd-indexed nodes are grouped together
// followed by all even-indexed nodes. The operation is performed in-place.
// Example: 1->2->3->4->5 becomes 1->3->5->2->4
func (t *ListNode) oddEvenList() {
	odds := t          // Pointer to the current odd node
	evens := odds.Next // Pointer to the current even node
	eventList := evens // Store the start of the even node sublist

	// Traverse and rearrange the list
	for evens != nil && evens.Next != nil {
		odds.Next = evens.Next // Link odd node to the next odd node
		odds = odds.Next       // Move odd pointer forward

		evens.Next = odds.Next // Link even node to the next even node
		evens = evens.Next     // Move even pointer forward
	}
	odds.Next = eventList // Attach even sublist after odd sublist
}
