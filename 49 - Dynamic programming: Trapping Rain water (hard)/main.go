package main

import (
	"fmt"
	"time"
)

// This program solves the "Trapping Rain Water" problem using dynamic programming.
// Given an array representing building heights, calculate how much water can be trapped
// between the buildings after raining.
//
// Dynamic Programming Approach:
// 1. For each building position, water trapped depends on:
//    - Maximum height to its left
//    - Maximum height to its right
//    - Current building height
// 2. Water at position i = min(maxLeft, maxRight) - height[i]
// 3. Use array to store maximum right heights, calculate left max on the fly
//
// Example:
// Heights: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6 (water units trapped between buildings)
//
// Time Complexity: O(n) - two passes through array
// Space Complexity: O(n) - for right max heights array

func main() {
	buildings := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 3}
	start := time.Now()
	result := trapWater(buildings)
	fmt.Printf("Rain water trapped: %v\n", result)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func trapWater(buildings []int) int {
	if len(buildings) == 0 {
		return 0
	}
	right := make([]int, len(buildings)) // Stores max height to right of each position
	water := 0

	// First pass: calculate maximum heights to the right
	aux := 0
	for i := len(buildings) - 1; i >= 0; i-- {
		aux = max(aux, buildings[i])
		right[i] = aux
	}

	// Second pass: calculate trapped water using left max and stored right max
	aux = 0 // Now used for left maximum
	for i, value := range buildings {
		aux = max(aux, value)
		water += min(aux, right[i]) - value
	}

	return water
}
