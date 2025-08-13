package main

// This program solves the Coin Change problem using dynamic programming.
// Given a set of coin denominations and a target amount, it finds the minimum
// number of coins needed to make up that amount.
//
// The approach uses recursion with division to optimize the search:
// 1. For each coin, calculate how many times it can be used (division)
// 2. Recursively solve for the remaining amount (modulo)
// 3. Track the minimum number of coins needed across all possibilities
//
// Time Complexity: O(n*m) where n is amount and m is number of coin denominations
// Space Complexity: O(n) due to recursion stack depth

import (
	"fmt"
	"math"
	"time"
)

// Change holds the available coin denominations
type Change struct {
	coins []int
}

func main() {
	change := Change{coins: []int{1, 2, 5, 10, 20, 50, 100}}

	start := time.Now()
	result := change.coinChange(103)
	elapsed := time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time elapsed from optimized approach: %s\n", elapsed)
}

// coinChange finds the minimum number of coins needed to make up the target amount
func (c *Change) coinChange(amount int) int {
	if amount == 0 {
		return 0
	}
	result := math.MaxInt32

	for _, coin := range c.coins {
		if amount < coin {
			continue
		}
		div := amount / coin
		res := c.coinChange(amount % coin)
		if (div + res) < result {
			result = div + res
		}
	}
	return result
}
