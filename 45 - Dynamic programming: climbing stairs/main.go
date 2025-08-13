package main

import (
	"fmt"
	"time"
)

// This program solves the climbing stairs problem using dynamic programming.
// Given n stairs, find the number of distinct ways to climb to the top,
// where you can either climb 1 or 2 steps at a time.
//
// Three approaches are implemented:
// 1. Top-down recursive approach (exponential time complexity)
// 2. Bottom-up dynamic programming approach (linear time and space complexity)
// 3. Bottom-up optimized approach (linear time, constant space)
//
// Example:
// Input: n = 3
// Output: 3
// Explanation: There are 3 ways:
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
//
// The problem follows the Fibonacci sequence pattern where:
// F(n) = F(n-1) + F(n-2)
// This is because to reach step n, we can either:
// - Take 1 step from n-1
// - Take 2 steps from n-2

func main() {
	n := 10
	start := time.Now()
	result := numWaysTop(n)
	elapsed := time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time elapsed from recursive approach: %s\n", elapsed)

	start = time.Now()
	result = numWaysBottom(n)
	elapsed = time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time elapsed from bottom-up approach: %s\n", elapsed)

	start = time.Now()
	result = numWaysOptimization(n)
	elapsed = time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time elapsed from optimized approach: %s\n", elapsed)
}

// numWaysTop calculates number of ways to climb stairs recursively (top-down)
// It recursively breaks down the problem into subproblems:
// ways(n) = ways(n-1) + ways(n-2)
// Base cases are n=0 (1 way) and n<0 (0 ways)
//
// Parameters:
//   - steps: number of remaining steps to climb
//
// Returns: number of distinct ways to climb the stairs
//
// Time Complexity: O(2^n) due to recursive branching
// Space Complexity: O(n) due to recursion stack depth
func numWaysTop(steps int) int {

	if steps == 0 {
		return 1
	}
	if (steps - 2 + steps - 1) <= 0 {
		return 1
	}
	return numWaysTop(steps-1) + numWaysTop(steps-2)
}

// numWaysBottom calculates number of ways using bottom-up dynamic programming
// Builds solution iteratively by storing results in array:
// dp[i] = dp[i-1] + dp[i-2]
//
// Parameters:
//   - steps: total number of steps to climb
//
// Returns: number of distinct ways to climb the stairs
//
// Time Complexity: O(n) where n is number of steps
// Space Complexity: O(n) for dynamic programming array
func numWaysBottom(steps int) int {
	ans := make([]int, steps+1)
	ans[0] = 1
	ans[1] = 1
	for i := 2; i <= steps; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}
	return ans[steps]
}

// numWaysOptimization calculates number of ways using optimized bottom-up approach
// Uses only two variables to track previous states instead of an array
// This reduces space complexity to O(1) while maintaining linear time
//
// Parameters:
//   - steps: total number of steps to climb
//
// Returns: number of distinct ways to climb the stairs
//
// Time Complexity: O(n) where n is number of steps
// Space Complexity: O(1) using only two variables
func numWaysOptimization(steps int) int {
	oneSteps := 1
	twoSteps := 1
	for i := 2; i <= steps; i++ {
		current := oneSteps + twoSteps
		oneSteps = twoSteps
		twoSteps = current
	}
	return twoSteps
}
