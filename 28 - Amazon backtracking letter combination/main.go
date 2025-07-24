package main

import (
	"fmt"
	"sync"
	"time"
)

type PhoneText struct {
	ans           []string // ans holds all generated letter combinations
	digitToString map[int]string
	mu            sync.Mutex
	wg            sync.WaitGroup
}

// main is the entry point of the application. It creates a Solution instance,
// calls letterCombination with a sample input, and prints the result.
// Eventhough we don't need go routines I've applied one just to show how it may work.
func main() {
	// Create a Solution instance
	s := PhoneText{
		digitToString: map[int]string{
			2: "abc",
			3: "def",
			4: "ghi",
			5: "jkl",
			6: "mno",
			7: "pqrs",
			8: "tuv",
			9: "wxyz",
		},
	}
	start := time.Now()
	// Call letterCombination with an anonymous function returning the digits [2, 3, 4]
	s.letterCombination(func() []int { return []int{2, 3, 4} }, "", 0, true)
	s.wg.Wait()
	// Print all generated combinations
	fmt.Printf("%v\n", s.ans)

	s.ans = []string{}
	s.letterCombination(func() []int { return []int{3, 4, 5, 7, 8, 9, 2, 6} }, "", 0, true)
	s.wg.Wait()
	//s.wg.Wait()
	// Print all generated combinations
	fmt.Printf("%v\n", s.ans)
	elapsed := time.Since(start)
	fmt.Printf("%v\n", elapsed)
}

// solution is a recursive helper method for generating all combinations via backtracking.
// It builds combinations by appending one letter at a time for each digit.
//
// Parameters:
//
//	digits: []int                - slice of input digits
//	digitToString: map[int]string - maps each digit to its corresponding string of letters
//	cur: string                  - current combination being built
//	digitIndex: int              - current index in the digits slice
func (s *PhoneText) letterCombination(digits func() []int, cur string, digitIndex int, isRoot bool) {
	// It's ok to check this in the backtraking method as digits are constant all the time.
	if len(digits()) == 0 {
		return // No digits to process
	}
	// Base case: if the current combination length matches the input digits, add to the answer list.
	if len(cur) == len(digits()) {
		// Always protect shared mutable data structures ([]string, map) with sync.Mutex or channels.
		s.mu.Lock()
		s.ans = append(s.ans, cur) // Add the current combination to the result slice.
		s.mu.Unlock()
		return // End this recursion branch.
	}
	// Get the current digit to process based on the current index.
	currentDigit := int(digits()[digitIndex])
	// Look up the possible letters for the current digit.
	currentString := s.digitToString[currentDigit]

	// Iterate over each possible letter for the current digit.
	for _, char := range currentString {
		// Recursively build the next combination by adding the current letter and moving to the next digit.
		next := cur + string(char)
		if isRoot && digitIndex == 0 {
			s.wg.Add(1)
			go func(n string) {
				defer s.wg.Done()
				s.letterCombination(digits, n, digitIndex+1, false)
			}(next)
		} else {
			s.letterCombination(digits, next, digitIndex+1, false)
		}
	}
	//s.wg.Done()
}
