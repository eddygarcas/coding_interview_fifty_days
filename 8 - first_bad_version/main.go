package main

import (
	"fmt"
	"time"
)

/*
You have N versions of a software, we want to find the first "bad" version of the software
if a version is bad, all versions after it are bad also
Will start from 1 to N same the list is ordered so you can do a binary search again
- First step would be go directly to the middle of the sorted array
- Check if the middle value is a bad version
- If it's a bad version means that further versions are also bad so move the right pointer pointing to middle
- Loop again while left pointer is still lower than right pointer
- In case the middle value is a valid version move the left point to be the middle value plus one, as now a
valid version could be further version between left and right pointer
- Eventually when the left pointer points to the same element as the right pointer and middle would be the first bad version.
*/
func main() {
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14}
	mid := len(e) / 2
	r := len(e)
	l := 0

	start := time.Now()
	for l < r {
		if isBadVersion(e[mid]) {
			r = mid
		} else {
			l = mid + 1
		}
		mid = (r + l) / 2
	}
	fmt.Println(e[mid])
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)

}

func isBadVersion(version int) bool {
	return version >= 11
}
