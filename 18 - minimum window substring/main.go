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
	result := minSubstring(strings.Split(s, ""), t)
	// Print the result.
	elapsed := time.Since(start)
	fmt.Printf("%s\n", result)
	fmt.Printf("%s\n", elapsed)
}

// minSubstring returns the minimum window substring from 's' (as a slice of characters)
// that contains all characters from string 't'. If no such window exists, returns an empty string.
func minSubstring(s []string, t string) string {
	// Initialize two pointers, 'l' and 'r', to the start of the string 's'.
	l := 0
	r := 0
	// Initialize the result string.
	result := ""
	// Iterate over the string 's' using the two pointers.
	for l < len(s) && r < len(s) {
		// Check if the current window 's[l:r]' contains all characters of 't'.
		if isSubet(s[l:r], strings.Split(t, "")) {
			// If it does, update the result string and move the left pointer to the right.
			result = strings.Join(s[l:r], "")
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
	set := make(map[string]int)
	// Initialize a counter for the number of unmatched characters.
	count := 0
	// Count the frequency of each character in 't'.
	for _, t := range t {
		set[t]++
	}
	// Iterate over the characters in 's'.
	for _, s := range s {
		// If the character is present in 't', increment count.
		if _, ok := set[s]; ok {
			count++
		}
	}
	// Eventually count has to be the same as the length of target array.
	return count == len(t)
}
