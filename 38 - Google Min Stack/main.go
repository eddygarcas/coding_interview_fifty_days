package main

import "fmt"

// Integer is a type constraint interface that accepts various integer types
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// stack represents a generic stack data structure that can store integer types
type stack[T Integer] struct {
	elems []T // slice to store stack elements
}

// push adds an element to the top of the stack
func (s *stack[T]) push(elems T) {
	s.elems = append(s.elems, elems)
}

// pop removes and returns the top element from the stack
// Returns error if stack is empty
func (s *stack[T]) pop() (T, error) {
	if len(s.elems) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	e := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return e, nil
}

// top returns the element at the top of the stack without removing it
func (s *stack[T]) top() T {
	return s.elems[len(s.elems)-1]
}

// getMin finds the minimum element in the stack while preserving stack order
func (s *stack[T]) getMin() T {
	aux := stack[T]{} // auxiliary stack for temporary storage
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

// order restores the original stack order using an auxiliary stack
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
	stack.push(-10)
	fmt.Printf("Top stack: %v\n", stack.top())
	fmt.Printf("Minimum element in stack: %v\n", stack.getMin())
	fmt.Printf("Original stack: %v\n", stack.elems)
	fmt.Printf("Top stack: %v\n", stack.top())
}
