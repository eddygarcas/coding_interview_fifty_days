package main

import (
	"fmt"
	"time"
)

// main is the entry point of the program. It initializes the input array, measures the execution time, and prints all subsets.
func main() {
	nums := []int{1, 2, 3}

	start := time.Now()
	subs := subsets(nums)
	elapsed := time.Since(start)
	fmt.Printf("Subset time: %v\n", elapsed)

	fmt.Printf("%v\n", subs)
}

// subsets returns all possible subsets (the power set) of the given integer slice nums.
// It uses a backtracking approach to generate all combinations.
func subsets(nums []int) [][]int {
	ans := make([][]int, 0) // Stores all subsets
	cur := make([]int, 0)   // Current subset being built
	recurSubset(&nums, &ans, cur, 0)

	return ans
}

// recurSubset is a recursive helper function for generating subsets.
// nums: pointer to the input slice
// ans: pointer to the result slice of slices
// cur: current subset being constructed
// index: current index in nums to consider
func recurSubset(nums *[]int, ans *[][]int, cur []int, index int) {
	if index > len(*nums) {
		return
	}
	*ans = append(*ans, append([]int(nil), cur...)) // Add a copy of the current subset

	for i := index; i < len(*nums); i++ {
		// Skip duplicates (not necessary for unique input, but useful for generalized case)
		if i > index && (*nums)[i] == (*nums)[i-1] {
			continue
		}
		cur = append(cur, (*nums)[i])
		recurSubset(nums, ans, cur, i+1)
		cur = cur[:len(cur)-1] // Backtrack
	}
}
