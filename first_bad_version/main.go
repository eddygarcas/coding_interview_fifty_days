package main

import "fmt"

// You have N versions of a software, we want to find the first "bad" version of the software
// if a version is bad, all versions after it are bad also
// Will start from 1 to N same the list is ordered so you can do a binary search again

func main() {
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14}
	mid := len(e) / 2
	l := 0
	for l < len(e) {
		if isBadVersion(e[mid]) {
			mid = mid / 2
		} else {
			l = mid + 1
			mid = (len(e)-l)/2 + mid
		}
	}
	fmt.Println(e[mid])

}

func isBadVersion(version int) bool {
	return version >= 10
}
