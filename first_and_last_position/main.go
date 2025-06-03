package main

import (
	"fmt"
	"time"
)

type Point struct {
	x, y int
}

func (p Point) toArray() []int {
	return []int{p.x, p.y}
}

func main() {
	e := []int{10, 13, 11, 15, 17, 18}
	target := 20
	result := Point{x: -1, y: -1}

	// Find first target occurrence left to right
	// Keep the index
	// from last index towards leftpointer find the last occurence
	// Could be the same
	start := time.Now()

	for index, element := range e {
		if element == target {
			if result.x == -1 {
				result = Point{x: index, y: index}
			} else {
				result.y = index
			}
		}
	}
	fmt.Printf("First and last position is %v\n", result.toArray())
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}
