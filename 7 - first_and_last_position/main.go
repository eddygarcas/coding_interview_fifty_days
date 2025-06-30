package main

import (
	"fmt"
	"time"
)

type point struct {
	x, y int
}

func (p point) ToArray() []int {
	return []int{p.x, p.y}
}

func main() {
	e := []int{10, 11, 11, 11, 12, 13, 18, 19, 20, 21, 22, 23, 24, 25, 26, 26}
	target := 19
	result := point{x: -1, y: -1}

	// Another solution would be once you've found the first one, iterate backwards, because in case of a very large
	// array doing it linearly may take long time.
	start := time.Now()
	for index, element := range e {
		if element == target {
			if result.x == -1 {
				result = point{x: index, y: index}
			} else {
				result.y = index
			}
		}
	}
	fmt.Printf("First and last position is %v\n", result.ToArray())
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)

	// Implement Binary search
	// Divide the array by two and look from 0 to mid
	start = time.Now()
	result.x = findFirst(e, target)
	result.y = findLast(e, target)
	fmt.Printf("First and last position is %v\n", result.ToArray())
	elapsed = time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}

func findFirst(e []int, target int) int {
	left, right := 0, len(e)-1
	for right >= left {
		mid := (right + left) / 2
		if e[mid] == target {
			if mid == 0 || e[mid-1] != target {
				return mid
			}
			right = mid - 1
		} else if e[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func findLast(e []int, target int) int {
	left, right := 0, len(e)-1
	for right >= left {
		mid := (right + left) / 2
		if e[mid] == target {
			if mid == len(e)-1 || e[mid+1] != target {
				return mid
			}
			left = mid + 1
		} else if e[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
