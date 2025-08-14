package main

// This program solves the Coin Change problem using dynamic programming.
// Given a set of coin denominations and a target amount, it finds the minimum
// number of coins needed to make up that amount.
//
// Example:
// coins = [1,2,5], amount = 11
// Output: 3 (11 = 5 + 5 + 1)
//
// Two approaches are implemented:
// 1. Top-down recursive with division optimization
// 2. Bottom-up dynamic programming with tabulation
//
// Time Complexity: O(n*m) where n is amount and m is number of coin denominations
// Space Complexity: O(n) for tabulation array

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
	result := change.coinChange(30)
	elapsed := time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time elapsed from optimized approach: %s\n", elapsed)

	start = time.Now()
	result = change.coinsChangeNested(30)
	elapsed = time.Since(start)
	fmt.Printf("Result recursive: %d\n", result)
	fmt.Printf("Time elapsed from recursive approach: %s\n", elapsed)
}

// coinChange finds minimum coins needed using recursive division approach
// Time Complexity: O(n*m) where n is amount and m is number of coins
// Space Complexity: O(n) due to recursion stack depth
//
// Algorithm steps:
//  1. Base case: if amount is 0, return 0 coins needed
//  2. For each coin denomination:
//     a. Skip if coin value exceeds remaining amount
//     b. Calculate maximum times coin can be used (division)
//     c. Recursively solve for remaining amount (modulo)
//     d. Update minimum coins needed if current solution is better
//  3. Return minimum coins found or MaxInt if no solution exists
func (c *Change) coinChange(amount int) int {
	if amount <= 0 {
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

// coinsChangeNested finds minimum coins using bottom-up dynamic programming
// Time Complexity: O(n*m) where n is amount and m is number of coins
// Space Complexity: O(n) for dp array
//
// Algorithm steps:
//  1. Create dp array of size amount+1, initialize with MaxInt
//  2. Set base case: dp[0] = 0 (zero coins needed for amount 0)
//  3. For each amount from 1 to target:
//     a. For each coin denomination:
//     - If coin <= current amount and dp[amount-coin] is valid
//     - Update dp[amount] with minimum of current or (dp[amount-coin] + 1)
//  4. Return dp[amount] if valid solution exists, else return -1
func (c *Change) coinsChangeNested(amount int) int {
	if amount <= 0 {
		return 0
	}
	result := make([]int, amount+1)
	for i := range result {
		result[i] = math.MaxInt32
	}
	result[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range c.coins {
			if coin <= i && result[i-coin] != math.MaxInt32 {
				result[i] = int(math.Min(float64(result[i-coin]+1), float64(result[i])))
			}
		}
		if result[amount] != math.MaxInt32 {
			return result[amount]
		}
	}
	return -1
}
