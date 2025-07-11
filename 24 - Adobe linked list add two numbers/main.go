package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	sum1 := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3}}}

	sum2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 4}}}

	fmt.Printf("Both sum %v and %v\n", sum1, sum2)

	result := addTwoNumber(sum1, sum2)

	for result != nil {
		fmt.Printf("%d,", result.Val)
		result = result.Next
	}
	fmt.Printf("Result %v", result)
}

func addTwoNumber(t, b *ListNode) *ListNode {
	var ret *ListNode = &ListNode{}
	ans := ret
	var carry int = 0
	for t != nil || b != nil {
		sum := (t.Val + b.Val + carry) % 10
		carry = (t.Val + b.Val) / 10
		fmt.Printf("Sum %d and Carry %d\n", sum, carry)
		ret.Next = &ListNode{Val: sum}
		ret = ret.Next
		t = t.Next
		b = b.Next
	}

	return ans
}
