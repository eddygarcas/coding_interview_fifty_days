package main

// This program solves the Best Time to Buy and Sell Stock problem using dynamic programming.
// Given an array of stock prices where prices[i] is the price on day i, find the maximum
// profit by choosing a single day to buy and a later day to sell.
//
// The algorithm uses a one-pass approach by tracking:
// - Minimum price seen so far (best buying opportunity)
// - Maximum profit possible by comparing (current price - minimum price)
// - Lower index to track the day with lowest price
// - Upper index to track the day with highest price after lowest price
//
// Example:
// Input prices: [1,5,3,0,9,4,1,15]
// Result: Buy at $1 (day 0), sell at $15 (day 7) for profit of $14
//
// Time Complexity: O(n) - single pass through prices array
// Space Complexity: O(1) - only tracking few variables regardless of input size

import (
	"errors"
	"fmt"
	"time"
)

// Stocks holds the price data and tracking indices for best buy/sell days
type Stocks struct {
	prices     []int // Array of daily stock prices
	lowerIndex int   // Index of day with lowest price (best buy day)
	upperIndex int   // Index of day with highest price after lowest price (best sell day)
}

func main() {
	stocks := Stocks{
		prices:     []int{7, 1, 5, 3}, // Example: Buy at $1, sell at $6 for profit of $5
		lowerIndex: -1,                // Initialize to invalid index
		upperIndex: -1,                // Initialize to invalid index
	}

	start := time.Now()
	minPrice := stocks.prices[0]
	maxProfit := 0
	_, profit, err := stocks.bestSellingPrice(minPrice, maxProfit, 0)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Best buy price: %d, Best sell price: %d\n", stocks.prices[stocks.lowerIndex], stocks.prices[stocks.upperIndex])
		fmt.Printf("Maximum profit: %d\n", profit)
	}
	fmt.Printf("Time elapsed: %s\n", time.Since(start))
}

// bestSellingPrice recursively finds the optimal buying and selling prices to maximize profit.
// The algorithm works by:
// 1. Base case: when we reach end of prices array (i >= len(prices))
//   - Check if we found valid buy/sell days (lowerIndex != -1 and before upperIndex)
//   - Return error if no valid combination found
//
// 2. For each price:
//   - Update minPrice if current price is lower
//   - If new minPrice found, update lowerIndex if it's sequential
//   - Calculate profit with current price (as potential sell price)
//   - If profit is higher than maxProfit, update maxProfit and upperIndex
//
// 3. Recursively process next price
//
// Parameters:
//
//	minPrice: lowest price seen so far
//	maxProfit: highest profit possible from prices seen so far
//	i: current index in prices array
//
// Returns:
//
//	minPrice: final minimum price found
//	maxProfit: maximum profit possible
//	error: if no valid buy/sell combination found
func (s *Stocks) bestSellingPrice(minPrice, maxProfit, i int) (int, int, error) {
	if i >= len(s.prices) {
		if s.lowerIndex == -1 || s.lowerIndex > s.upperIndex {
			return -1, -1, errors.New("No valid buy price")
		}
		return minPrice, maxProfit, nil
	}

	minPrice = min(minPrice, s.prices[i])
	if minPrice == s.prices[i] && i <= s.lowerIndex+1 {
		s.lowerIndex = i
	}

	currentProfit := s.prices[i] - minPrice
	if currentProfit > maxProfit {
		maxProfit = currentProfit
		s.upperIndex = i
	}

	return s.bestSellingPrice(minPrice, maxProfit, i+1)
}
