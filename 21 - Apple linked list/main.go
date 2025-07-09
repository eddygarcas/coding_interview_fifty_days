// Package main implements a solution for merging two sorted linked lists.
// This is a common interview problem that tests understanding of linked list operations.
package main

import "fmt"

// ListNode represents a node in a singly linked list.
// Each node contains an integer value and a pointer to the next node.
type ListNode struct {
	Val  int       // The value stored in this node
	Next *ListNode // Pointer to the next node in the list, or nil if this is the last node
}

// main demonstrates the merging of two sorted linked lists.
// The algorithm maintains the sorted order in the resulting merged list.
func main() {
	// Initialize a dummy head node for the result list
	// This simplifies the merging process by avoiding special cases for the head
	cur := &ListNode{}
	ans := cur // Save the dummy head to traverse the final result

	// Create two sorted linked lists for demonstration
	// list1: 1 -> 2 -> 3 -> 4
	var list1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	// list2: 1 -> 6 -> 7 -> 8
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8}}}}

	// Merge the two lists by comparing values and linking nodes in sorted order
	// Continue until we've processed all nodes from both lists
	for list1 != nil || list2 != nil {
		// If list1 is empty or list2's current value is smaller, take from list2
		if list1 == nil || (list2 != nil && list1.Val > list2.Val) {
			cur.Next = list2
			list2 = list2.Next
		} else {
			// Otherwise, take from list1
			cur.Next = list1
			list1 = list1.Next
		}
		// Move the current pointer forward in the result list
		cur = cur.Next
	}

	// Print the merged list (skipping the dummy head node)
	// The result should be a sorted list: 1 -> 1 -> 2 -> 3 -> 4 -> 6 -> 7 -> 8
	for ans.Next != nil {
		fmt.Println(ans.Next.Val)
		ans = ans.Next
	}
}
