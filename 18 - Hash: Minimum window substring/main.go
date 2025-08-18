package main

import (
	"fmt"
	"strings"
	"time"
)

// main demonstrates finding the minimum window substring in 's' containing all characters of 't'.
func main() {
	// Define the string 's' and the substring 't' to search for.
	s := "abaaccadaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	t := "aca"
	start := time.Now()
	// Find the minimum window substring in 's' containing all characters of 't'.
	result := minSubstring(strings.Split(s, ""), strings.Split(t, ""))
	// Print the result.
	elapsed := time.Since(start)
	fmt.Printf("%s\n", result)
	fmt.Printf("%s\n", elapsed)
}

// minSubstring returns the minimum window substring from 's' (as a slice of characters)
// that contains all characters from string 't'. If no such window exists, returns an empty string.
func minSubstring(s, t []string) []string {
	// Initialize two pointers, 'l' and 'r', to the start of the string 's'.
	l := 0
	r := 0
	// Initialize the result as array of strings
	var result []string
	// Iterate over the string 's' using the two pointers.
	for l < len(s) && r < len(s) {
		// Check if the current window 's[l:r]' contains all characters of 't'.
		if isSubet(s[l:r], t) {
			// If it does, update the result string and move the left pointer to the right.
			if len(result) == 0 || len(s[l:r]) < len(result) {
				result = s[l:r]
			} else if len(result) == len(t) {
				return result
			}
			l++
		} else {
			// If it doesn't, move the right pointer to the right to expand the window.
			r++
		}
	}
	// Return the result string.
	return result
}

// isSubet checks if all characters in t are present in s (with at least the same frequency).
// Both s and t are slices of single-character strings.
func isSubet(s, t []string) bool {
	// Create a frequency map for the characters in 't'.
	set := make(map[string]uint8, len(t))
	// Initialize a counter for the number of unmatched characters.
	result := 0
	// Count the frequency of each character in 't'.
	for _, t := range t {
		set[t]++
	}
	// Iterate over the characters in 's'.
	for _, s := range s {
		// If the character is present in 't', increment count.
		if val, ok := set[s]; ok && val > 0 {
			set[s]--
			result++
		}
	}
	// Eventually count has to be the same as the length of target array.
	return result == len(t)
}
