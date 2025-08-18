package main

// This program finds the length of the longest substring without repeating characters.
// It uses a sliding window approach and a map to track the last seen index of each character.

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	// e is the input byte slice to analyze
	e := []byte{'a', 'c', 'c', 'd', 'e'}
	// subsMap stores the last seen index for each character
	subsMap := make(map[byte]int, len(e))

	ans := 0   // Stores the length of the longest substring found
	left := 0  // Left boundary of the current window
	right := 0 // Right boundary of the current window
	for i := 0; i < len(e); i++ {
		if val, ok := subsMap[e[i]]; ok {
			// If character was seen, move left boundary right after last occurrence
			left = int(math.Max(float64(left), float64(val+1)))
		}
		right++
		subsMap[e[i]] = i // Update last seen index
		ans = int(math.Max(float64(right-left), float64(ans)))
	}
	fmt.Printf("Max substring: %d\n", ans)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}
