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
// For input "baba", the isPalindrome matrix will look like:
//
//    b  a  b  a
// b  T  F  T  F
// a  -  T  F  T
// b  -  -  T  F
// a  -  -  -  T
//
// where T = true (is palindrome), F = false, - = unused
// Reading the matrix shows palindromes: "b","a","b","a" (len 1)
//                                      "bab" (len 3)
// Therefore longest palindrome is "bab"
//
// Time Complexity: O(n²) where n is string length (nested loops)
// Space Complexity: O(n²) for the 2D DP array

func main() {
	start := time.Now()
	result := longestPalindrome("baba")
	elapsed := time.Since(start)
	fmt.Printf("Longest palindrome: %v\n", result)
	fmt.Printf("Time taken: %v\n", elapsed)
}

// longestPalindrome finds the longest palindromic substring in string s
// Uses bottom-up dynamic programming with 2D boolean array
// Returns the longest palindrome found
// For input "baba", the isPalindrome matrix will look like:
//
//	b  a  b  a
//
// b  T  F  T  F
// a  -  T  F  T
// b  -  -  T  F
// a  -  -  -  T
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
