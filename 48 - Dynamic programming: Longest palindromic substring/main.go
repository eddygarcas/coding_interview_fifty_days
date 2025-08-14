package main

import (
	"fmt"
	"time"
)

// This program finds the longest palindromic substring in a given string using dynamic programming.
// It uses a 2D boolean array to track palindrome status of all substrings.
//
// Dynamic Programming Approach:
// 1. A substring is palindrome if:
//    - First and last characters match
//    - Inner substring is palindrome (or length <= 1)
// 2. Fill DP table bottom-up, checking all possible substring lengths
// 3. Track longest palindrome found so far
//
// Time Complexity: O(n²) where n is string length (nested loops)
// Space Complexity: O(n²) for the 2D DP array

func main() {
	start := time.Now()
	result := longestPalindrome("baba")
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %v\n", elapsed)
	fmt.Printf("Palindrome: %v\n", result)
}

// longestPalindrome finds the longest palindromic substring in string s
// Uses bottom-up dynamic programming with 2D boolean array
// Returns the longest palindrome found
func longestPalindrome(s string) string {
	ans := ""
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	// Single characters are palindromes
	for i := 0; i < n; i++ {
		isPalindrome[i][i] = true
	}

	// Check all possible substrings
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			middleStringLen := i - j - 1
			isSameChar := s[i] == s[j]

			if isSameChar && (middleStringLen <= 1 || isPalindrome[j+1][i-1]) {
				isPalindrome[j][i] = true
				if i-j >= len(ans) {
					ans = s[j : i+1]
				}
			}
		}
	}
	return ans
}
