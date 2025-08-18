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
	// Input: "accde" visualized step by step:
	//
	// Step 1: [a]ccde    left=0, right=1, ans=1
	//          ^
	//
	// Step 2: [ac]cde    left=0, right=2, ans=2
	//          ^^
	//
	// Step 3: a[c]cde    left=1, right=3, ans=2
	//           ^        (duplicate 'c' found, move left)
	//
	// Step 4: ac[cde]    left=2, right=5, ans=3
	//             ^^^    (final window with longest unique substring)
	e := []byte{'a', 'c', 'c', 'd', 'e'}
	subsMap := make(map[byte]int, len(e))
	ans := 0
	left := 0
	right := 0

	for i := 0; i < len(e); i++ {
		if val, ok := subsMap[e[i]]; ok {
			left = int(math.Max(float64(left), float64(val+1)))
		}
		right++
		subsMap[e[i]] = i
		ans = int(math.Max(float64(right-left), float64(ans)))
	}

	fmt.Printf("Max substring: %d\n", ans)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}
