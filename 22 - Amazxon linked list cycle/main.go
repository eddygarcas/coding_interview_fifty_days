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

// main demonstrates the Floyd's cycle detection algorithm (Tortoise and Hare)
// by creating a linked list with a cycle and checking for a cycle's entry point.
// Also that recursive functions are way faster than tha loop to resolve this challenge
func main() {
	node1 := &ListNode{Val: 9}
	node2 := &ListNode{Val: 7}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 8}
	node6 := &ListNode{Val: 1}

	// Create a cycle: node6.Next points back to node3
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node3

	t := node1 // Tortoise pointer
	h := node1 // Hare pointer

	start := time.Now()
	if result, ok := floydAlg(t, h); ok {
		fmt.Printf("result: %v\n", result)
	} else {
		fmt.Printf("not ok")
	}
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %s\n", elapsed)

	start = time.Now()
	if result, ok := floydAlgRecur(t, h.Next.Next); ok {
		fmt.Printf("result: %v\n", result)
	} else {
		fmt.Printf("not ok")
	}
	elapsed = time.Since(start)
	fmt.Printf("elapsed: %s\n", elapsed)
}

// floydAlg implements Floyd's cycle detection algorithm (Tortoise and Hare).
// It takes two pointers (tortoise and hare) and checks for a cycle in the linked list.
// Returns the value at the node where the cycle is detected and true if a cycle exists,
// otherwise returns 0 and false.
func floydAlg(t, h *ListNode) (int, bool) {
	for h != nil && h.Next != nil {
		fmt.Printf("t val %d and h = %d\n", t.Val, h.Val)
		h = h.Next.Next
		t = t.Next
		if t == h {
			return h.Val, true
		}
	}
	return 0, false
}

// floydAlg implements Floyd's cycle detection algorithm (Tortoise and Hare).
// It takes two pointers (tortoise and hare) and checks for a cycle in the linked list.
// Returns the value at the node where the cycle is detected and true if a cycle exists,
// otherwise returns 0 and false.
// This time will implement a recursive method to make it even faster than a loop
func floydAlgRecur(t, h *ListNode) (int, bool) {
	fmt.Printf("t val %d and h = %d\n", t.Val, h.Val)
	if t == h {
		return h.Val, true
	} else if h.Next == nil {
		return 0, false
	}
	return floydAlgRecur(t.Next, h.Next.Next)
}
