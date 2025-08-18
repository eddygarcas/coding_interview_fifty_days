package main

import (
	"fmt"
	"strings"
	"time"
)

// This program solves the Minimum Window Substring problem using a sliding window and hash map approach.
// Given two strings 's' and 't', find the minimum window in 's' that contains all characters of 't'.
//
// Example:
// s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC" (smallest substring containing all characters from t)
//
// Algorithm:
// 1. Use two pointers l and r to define a window
// 2. Expand window by moving r until all characters of t are found
// 3. Contract window by moving l while still maintaining all required characters
// 4. Track minimum valid window found so far
//
// Time Complexity: O(n * m) where n is length of s and m is length of t
// Space Complexity: O(k) where k is size of character set (constant for ASCII)

// main demonstrates finding the minimum window substring in 's' containing all characters of 't'.
func main() {
	s := "abaaccadaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	t := "aca"
	start := time.Now()
	result := minSubstring(strings.Split(s, ""), strings.Split(t, ""))
	elapsed := time.Since(start)
	fmt.Printf("%s\n", result)
	fmt.Printf("%s\n", elapsed)
}

// minSubstring returns the minimum window substring from 's' that contains all characters from 't'.
// Uses sliding window technique with two pointers to find minimum valid substring.
// Time Complexity: O(n * m) where n is length of s and m is length of t
// Space Complexity: O(1) for result slice
func minSubstring(s, t []string) []string {
	l := 0
	r := 0
	var result []string
	for l < len(s) && r < len(s) {
		if isSubet(s[l:r], t) {
			if len(result) == 0 || len(s[l:r]) < len(result) {
				result = s[l:r]
			} else if len(result) == len(t) {
				return result
			}
			l++
		} else {
			r++
		}
	}
	return result
}

// isSubet checks if all characters in t are present in s with required frequency.
// Uses a hash map to track character frequencies.
// Time Complexity: O(m) where m is length of input slice s
// Space Complexity: O(k) where k is size of character set
func isSubet(s, t []string) bool {
	set := make(map[string]uint8, len(t))
	result := 0
	for _, t := range t {
		set[t]++
	}
	for _, s := range s {
		if val, ok := set[s]; ok && val > 0 {
			set[s]--
			result++
		}
	}
	return result == len(t)
}
