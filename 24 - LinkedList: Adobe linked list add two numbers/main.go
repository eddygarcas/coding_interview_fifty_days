package main

import (
	"fmt"
)

// ListNode defines a node in a singly linked list.
type ListNode struct {
	Val  int       // Value held by the node
	Next *ListNode // Pointer to the next node
}

func main() {
	// Create first number: 3 -> 5 -> 3
	sum1 := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3}}}

	// Create second number: 4 -> 5 -> 4
	sum2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 4}}}

	fmt.Printf("Both sum %v and %v\n", sum1, sum2)

	// Add the two numbers represented by the linked lists
	result := addTwoNumber(sum1, sum2)

	// Print the result as a linked list
	for result != nil {
		fmt.Printf("%d,", result.Val)
		result = result.Next
	}
}

// addTwoNumber adds two numbers represented by linked lists.
// Each node contains a single digit. Digits are stored in reverse order.
// Returns the sum as a new linked list in the same format.
func addTwoNumber(t, b *ListNode) *ListNode {
	ret := &ListNode{Val: 0} // Dummy head node for the result list
	ans := ret               // Pointer to build the result list
	carry := 0               // Carry for sums >= 10
	sum := 0
	_t := t // Pointer for first list
	_b := b // Pointer for second list
	for _t != nil || _b != nil {
		sum = carry
		if _t != nil {
			sum += _t.Val
			_t = _t.Next
		}
		if _b != nil {
			sum += _b.Val
			_b = _b.Next
		}
		carry = sum / 10
		ans.Next = &ListNode{Val: sum % 10}
		ans = ans.Next
	}
	// If there is a remaining carry, add it as a new node
	if carry > 0 {
		ans.Next = &ListNode{Val: carry}
	}

	return ret.Next // Return the next node after dummy head
}
