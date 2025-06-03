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
	e := []int{10, 11, 13, 14, 11, 18}
	target := 11
	result := point{x: -1, y: -1}

	// Find first target occurrence left to right
	// Keep the index
	// from last index towards leftpointer find the last occurence
	// Could be the same
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
}
