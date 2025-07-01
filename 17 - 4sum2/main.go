package main

// This program solves the 4Sum II problem: Given four integer arrays A, B, C, and D,
// compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] == 0.
// It uses a hash map to optimize the search for complement sums.

import "fmt"

func main() {
	// Example input arrays
	a := []int{1, 2}
	b := []int{-2, -1}
	c := []int{-1, 2}
	d := []int{0, 2}

	// result stores the sum of all pairs from a and b, and their counts
	result := make(map[int]int)

	// Populate the hash map with all possible sums of a and b
	ans := fourSumTwo(a, b, result, false)
	fmt.Println(ans) // Should print 0, as no checking is done in this call

	// For each possible sum of c and d, check if its negation exists in result
	ans = fourSumTwo(c, d, result, true)
	fmt.Println(ans) // Prints the count of tuples where the total sum is zero

	fmt.Printf("fourSumTwo(%d,%d)=%d\n", a, b, result)
}

// fourSumTwo either populates the result hash map with sums (if check is false)
// or counts how many pairs from a and b sum with pairs from c and d to zero (if check is true).
func fourSumTwo(a []int, b []int, result map[int]int, check bool) int {
	ans := 0
	for _, val := range a {
		for _, valb := range b {
			if check {
				// If check is true, look for the negative sum in the result map
				if _, ok := result[-(val + valb)]; ok {
					ans++
				}
			} else {
				// Otherwise, store the sum in the result map
				result[val+valb]++
			}
		}
	}
	return ans
}
