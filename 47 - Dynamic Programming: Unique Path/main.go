package main

import (
	"fmt"
	"time"
)

// This program solves the Unique Paths problem using dynamic programming.
// Given a m x n grid, starting from top-left corner, find number of unique paths
// to reach bottom-right corner. Can only move right or down at any point.
//
// Dynamic Programming Approach:
// 1. Create m x n grid and initialize first row and column to 1
//    (only one way to reach any cell in first row/column)
// 2. For each cell (i,j), paths = paths from above + paths from left
//    dp[i][j] = dp[i-1][j] + dp[i][j-1]
//
// Time Complexity: O(m*n) to fill the dp grid
// Space Complexity: O(m*n) for the dp grid

func main() {
	start := time.Now()
	matrix := make([][]int, 30)
	result := waysToBottom(matrix)
	elapsed := time.Since(start)
	fmt.Printf("Fixed grid result: %d\n", result)
	fmt.Printf("Time taken: %v\n", elapsed)

	start = time.Now()
	result = waysToBottomNonFixed(30, 30)
	elapsed = time.Since(start)
	fmt.Printf("Variable grid result: %d\n", result)
	fmt.Printf("Time taken: %v\n", elapsed)
}

// waysToBottom calculates number of unique paths in a fixed square grid
// Uses bottom-up dynamic programming approach with 2D array
// Parameters:
//   - matrix: square grid to calculate paths
//
// Returns: total number of unique paths to bottom-right
func waysToBottom(matrix [][]int) int {
	n := cap(matrix)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		matrix[i][0] = 1 // Initialize first column
		matrix[0][i] = 1 // Initialize first row
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}
	return matrix[n-1][n-1]
}

// waysToBottomNonFixed calculates paths for variable sized grid
// Uses bottom-up dynamic programming with 2D array
// Parameters:
//   - x: number of rows
//   - y: number of columns
//
// Returns: total number of unique paths to bottom-right
func waysToBottomNonFixed(x, y int) int {
	matrix := make([][]int, x)
	for i := 0; i < x; i++ {
		matrix[i] = make([]int, y)
		matrix[i][0] = 1 // Initialize first column
	}
	for j := 0; j < y; j++ {
		matrix[0][j] = 1 // Initialize first row
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}
	return matrix[x-1][y-1]
}
