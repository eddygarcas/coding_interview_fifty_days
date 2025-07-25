package main

import (
	"fmt"
	"time"
)

// palindrome struct stores the result of palindrome partitions
type palindrome struct {
	ans [][]string // stores all valid palindrome partitions
}

// main demonstrates palindrome partitioning by finding all possible partitions
// of a string where each substring is a palindrome
func main() {
	w := palindrome{}
	start := time.Now()
	w.solution("aabedltegnuetdsasaab", []string{})
	fmt.Printf("Palindrome: %v\n", w.ans)
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s\n", elapsed)
}

// solution implements backtracking to find all palindrome partitions
// Parameters:
//   - s: remaining string to partition
//   - curArr: current partition being built
func (w *palindrome) solution(s string, partition []string) {
	// Base case: if string is empty, we found a valid partition
	if len(s) == 0 {
		temp := make([]string, len(partition))
		copy(temp, partition)
		w.ans = append(w.ans, temp)
		return
	}

	// Try all possible prefixes of the remaining string
	// s = "aab"	partition = [a] 		i = 1
	// s = "ab" 	partition = [a,a] 		i = 1
	// s = "b" 		partition = [a,a,b] 	i = 1
	// s = ""		partition = [a,a,b] 	i = 1
	// s = "aab" 	partition = []			i = 2
	// s = "aab" 	partition = [a,a]		i = 2
	// s = "b" 		partition = [a,a]		i = 1
	// s = "b" 		partition = [a,a][b]	i = 1
	for i := 1; i <= len(s); i++ {
		curStr := s[0:i]
		if w.isPalindrome(curStr) {
			partition = append(partition, curStr)    // Add palindrome
			w.solution(s[i:], partition)             // Recurse with remaining string
			partition = partition[:len(partition)-1] // Backtrack
		}
	}
}

// isPalindrome checks if a string is a palindrome using two pointers
// Returns true if the string reads the same forwards and backwards
func (w *palindrome) isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
