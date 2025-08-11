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
	prospects := Prospects{houses: []int{1, 5, 3, 0, 9, 4}}
	route := make([]int, len(prospects.houses))

	start := time.Now()
	result := prospects.RobHouseFromBottom(len(prospects.houses)-1, route)
	fmt.Printf("result %d\n", result)
	elapsed := time.Since(start)
	fmt.Printf("elapsed %s\n", elapsed)
}

// RobHouseFromBottom calculates the maximum money that can be robbed from houses up to given address
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
func (h *Prospects) RobHouseFromBottom(address int, route []int) int {
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
	iStoleHouse := h.houses[address] + h.RobHouseFromBottom(address-2, route)
	iNotStoleHouse := h.RobHouseFromBottom(address-1, route)
	route[address] = max(iStoleHouse, iNotStoleHouse)
	return route[address]
}
