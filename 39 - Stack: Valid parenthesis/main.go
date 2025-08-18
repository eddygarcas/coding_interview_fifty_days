package main

import (
	"fmt"
	"time"
)

// Elems is a type constraint interface that accepts various numeric types
type Elems interface {
	~byte | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// stack represents a generic stack data structure that validates matching pairs
type stack[T Elems] struct {
	elems []T     // Slice to store stack elements
	match map[T]T // Map storing valid pairs of matching elements
}

// push adds an element to the top of the stack
func (s *stack[T]) push(elem T) {
	s.elems = append(s.elems, elem)
}

// pop removes and returns the top element from the stack
func (s *stack[T]) pop() T {
	if len(s.elems) == 0 {
		return 0
	}
	e := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return e
}

// top returns the element at the top of the stack without removing it
func (s *stack[T]) top() (T, error) {
	if len(s.elems) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.elems[len(s.elems)-1], nil
}

// isValid checks if the input sequence contains valid matching pairs
// Returns true if all elements match correctly according to the match map
func (s *stack[T]) isValid(input []T) bool {
	for _, digit := range input {
		if _, ok := s.match[digit]; ok {
			s.push(digit)
		} else {
			if s.match[s.pop()] != digit {
				return false
			}
		}
	}
	return true
}

func main() {
	var stack stack[byte]
	// Initialize match map for parentheses validation
	stack.match = map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	start := time.Now()
	input := []byte{'{', '(', '[', ']', ')', '}'}
	valid := stack.isValid(input)
	fmt.Printf("stack is valid? %t\n", valid)

	// Test with numeric pairs
	stack.match = map[byte]byte{
		0: 0,
		1: 1,
		2: 2,
	}
	input = []byte{0, 1, 2, 2, 1, 0, 4}
	valid = stack.isValid(input)
	fmt.Printf("stack is valid? %t\n", valid)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
