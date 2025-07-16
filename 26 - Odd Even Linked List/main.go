package main

import "fmt"

// ListNode defines a node in a singly linked list.
type ListNode struct {
	Val  int       // Value held by the node
	Next *ListNode // Pointer to the next node in the list
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

	// Apply the oddEvenList operation to rearrange the nodes
	result := node1
	node1.oddEvenList()
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}

// oddEvenList rearranges nodes in the list so that all odd-indexed nodes
// are grouped together followed by the even-indexed nodes. The first node is considered odd.
// The operation is performed in-place.
func (t *ListNode) oddEvenList() {
	odd := t             // Pointer to the current odd node
	even := odd.Next     // Pointer to the current even node
	evenList := odd.Next // Head of the even node sublist

	// Traverse and rearrange the list
	for even != nil && evenList.Next != nil {
		odd.Next = even.Next // Link current odd node to the next odd node
		odd = odd.Next       // Move odd pointer forward
		even.Next = odd.Next // Link current even node to the next even node
		even = even.Next     // Move even pointer forward
	}
	odd.Next = evenList // Attach even sublist after odd sublist
}
