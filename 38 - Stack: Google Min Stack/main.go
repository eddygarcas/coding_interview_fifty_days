package main

import (
	"fmt"
	"time"
)

// Integer is a type constraint interface that accepts various integer types
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// stack represents a generic stack data structure that can store integer types.
// It maintains both the elements and their corresponding minimum values.
type stack[T Integer] struct {
	elems [][]T // Each element is a pair [value, minSoFar]
}

// push adds an element to the top of the stack along with the minimum element up to that point.
// The minimum is stored alongside each element for O(1) min retrieval.
func (s *stack[T]) push(elem T) {
	var minElem T
	if s.top() == nil {
		minElem = elem
	} else {
		minElem = s.top()[1]
	}
	s.elems = append(s.elems, []T{elem, min(minElem, elem)})
}

// pop removes and returns the top element from the stack.
// Returns error if stack is empty.
func (s *stack[T]) pop() (T, error) {
	if len(s.elems) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	e := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return e[0], nil
}

// top returns the element at the top of the stack without removing it.
// Returns the value-minimum pair for the top element.
func (s *stack[T]) top() []T {
	if len(s.elems) == 0 {
		return nil
	}
	return s.elems[len(s.elems)-1]
}

// getMin finds the minimum element in the stack while preserving stack order.
// Uses an auxiliary stack to maintain original order.
// This method is deprecated as it increase the time complexity as it's not a linear approach
// @deprecated
func (s *stack[T]) getMin() T {
	aux := stack[T]{}
	var minElem T
	topElem, err := s.pop()
	minElem = topElem
	for err == nil {
		minElem = min(topElem, minElem)
		aux.push(topElem)
		topElem, err = s.pop()
	}
	s.order(aux)
	return minElem
}

// getMinPush returns the minimum element up to the top of the stack in O(1) time.
// Uses the pre-computed minimum values stored with each element.
func (s *stack[T]) getMinPush() T {
	return s.top()[1]
}

// order restores the original stack order using an auxiliary stack.
// Pops elements from aux stack and pushes them back to original stack.
func (s *stack[T]) order(aux stack[T]) {
	topElem, err := aux.pop()
	for err == nil {
		s.push(topElem)
		topElem, err = aux.pop()
	}
}

func main() {
	var stack stack[int32]
	stack.push(1)
	stack.push(-20)
	stack.push(3)
	stack.push(4)
	stack.push(-40)
	stack.push(-10)
	start := time.Now()
	fmt.Printf("Top stack: %v\n", stack.top())
	fmt.Printf("Minimum element in stack from push: %v\n", stack.getMinPush())
	fmt.Printf("Original stack: %v\n", stack.elems)
	fmt.Printf("Top stack: %v\n", stack.top())
	stack.pop()
	stack.pop()
	fmt.Printf("Minimum element in stack from push: %v\n", stack.getMinPush())
	fmt.Printf("Original stack: %v\n", stack.elems)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
