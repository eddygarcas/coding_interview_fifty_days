package main

import (
	"fmt"
	"time"
)

// This program solves the House Robber problem using dynamic programming.
// Problem: Given an array representing houses with money, determine maximum money
// that can be robbed without robbing adjacent houses.
//
// Dynamic Programming Approach:
// 1. For each house at index i, we have two choices:
//    a. Rob current house (i) and add it to max money from house (i-2)
//    b. Skip current house and take max money from house (i-1)
// 2. Use memoization array (route) to store computed results
// 3. Recurrence relation: dp[i] = max(nums[i] + dp[i-2], dp[i-1])
//
// Time Complexity: O(n) where n is number of houses
// Space Complexity: O(n) for memoization array

type Prospects struct {
	houses []int
}

func main() {
	prospects := Prospects{houses: []int{1, 5, 3, 0, 9, 4, 1, 15}}
	route := make([]int, len(prospects.houses))

	start := time.Now()
	result := prospects.RobHouseFromTop(len(prospects.houses)-1, route)
	fmt.Printf("result %d\n", result)
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed from the top %s\n", elapsed)

	start = time.Now()
	result = prospects.RobHouseFromBottom()
	fmt.Printf("result %d\n", result)
	elapsed = time.Since(start)
	fmt.Printf("Time elapsed from the bottom %s\n", elapsed)
}

// RobHouseFromTop calculates the maximum money that can be robbed from houses up to given address
// using dynamic programming with memoization.
// Parameters:
//   - address: current house index being considered
//   - route: memoization array storing max money for each house position
//
// The function achieves O(n) complexity by:
// 1. First checking if result for current address is already calculated (route[address] > 0)
// 2. If found, returns cached result instead of recalculating
// 3. Otherwise calculates optimal money by either:
//   - Including current house + max money from house (address-2)
//   - Excluding current house and taking max money from house (address-1)
//
// 4. Caches and returns result
func (h *Prospects) RobHouseFromTop(address int, route []int) int {
	if route[address] > 0 {
		return route[address]
	}
	if address < 0 {
		return 0
	}
	if address == 0 {
		return h.houses[0]
	}
	if address == 1 {
		return max(h.houses[0], h.houses[1])
	}
	iStoleHouse := h.houses[address] + h.RobHouseFromTop(address-2, route)
	iNotStoleHouse := h.RobHouseFromTop(address-1, route)
	route[address] = max(iStoleHouse, iNotStoleHouse)
	return route[address]
}

// RobHouseFromBottom calculates maximum money that can be robbed using bottom-up dynamic programming.
// It iteratively builds the solution by:
// 1. Handling base cases for 1-2 houses
// 2. For each house i, choosing maximum between:
//   - Previous max (not robbing current house)
//   - Current house value + max money from i-2 houses
//
// Parameters: none
// Returns: Maximum money that can be robbed
func (h *Prospects) RobHouseFromBottom() int {
	total := len(h.houses)
	if total == 1 {
		return h.houses[0]
	}
	// For houses: [1, 5, 3, 0, 9, 4, 1, 15]
	firstHouse := h.houses[0]                   // first = 1
	secondHouse := max(firstHouse, h.houses[1]) // second = max(1,5) = 5
	result := secondHouse                       // result = 5

	// i=2: house=3  -> result=max(5, 3+1=4)    -> first=5, second=5
	// i=3: house=0  -> result=max(5, 0+5=5)    -> first=5, second=5
	// i=4: house=9  -> result=max(5, 9+5=14)   -> first=5, second=14
	// i=5: house=4  -> result=max(14, 4+5=9)   -> first=14, second=14
	// i=6: house=1  -> result=max(14, 1+14=15) -> first=14, second=15
	// i=7: house=15 -> result=max(15, 15+14=29)-> first=15, second=29
	for i := 2; i < total; i++ {
		result = max(secondHouse, h.houses[i]+firstHouse)
		firstHouse = secondHouse
		secondHouse = result
	}
	return result
}
