package main

import (
	"fmt"
	"time"
)

// This program solves the "Combination Sum" problem using backtracking.
// Given a set of candidate numbers and a target sum, it finds all unique combinations
// of candidates where the chosen numbers sum to the target.
// Each number in candidates may be used an unlimited number of times.
//
// Example:
// Input: candidates = [2,3,6], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,3,2],[6,2]]

// candidates represents the data structure to store combinations and input numbers
type candidates struct {
	ans        [][]int // Stores all valid combinations
	candidates []int   // Input array of candidate numbers
	// cur []int
}

// main is the entry point - initializes data and runs the combination sum algorithm
func main() {
	start := time.Now()
	var cur []int
	s := candidates{make([][]int, 0), make([]int, 0)}
	s.solution(func() []int { return []int{2, 3, 6} }, &cur, 7, 0, 0)
	fmt.Printf("%v\n", s.ans)
	elapsed := time.Since(start)
	fmt.Printf("%v\n", elapsed)
}

// solution implements the backtracking algorithm to find all valid combinations
// Parameters:
//   - candidates: Function returning array of candidate numbers
//   - cur: Current combination being built
//   - target: Target sum to achieve
//   - index: Current index in candidates array
//   - sum: Running sum of current combination
//
// You can actually move cur []int to be part of the struct this way would avoid the pointer.
func (s *candidates) solution(candidates func() []int, cur *[]int, target int, index int, sum int) {
	if sum == target {
		// Found valid combination, add copy to results
		s.ans = append(s.ans, append([]int{}, *cur...))
	} else if sum < target {
		// Try adding each remaining candidate
		for i, candidate := range candidates()[index:] {
			*cur = append(*cur, candidate)                              // Add candidate
			s.solution(candidates, cur, target, i, sum+candidates()[i]) // Recurse
			// Will execute the following sentence for instance in this situation:
			// [2,2,2,2] -> because the previous execution SUM was be 8
			// So will pop cur to [2,2,2] and will run the next loop iteration
			*cur = (*cur)[:len(*cur)-1] // Backtrack
		}
	}
}
