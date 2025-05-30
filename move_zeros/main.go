package main

import (
	"fmt"
	"time"
)

func main() {
	elements := []int{6, 1, 0, 3, 0, 1, 0, 12}
	// Go through the array
	// If you find a 0 place it directly at the end of the array
	// Move the rest of the elements to the left until you read the index
	//Do this until the end of the array
	// This could be O(n2) would be better to first move all the non-zeros and then fill the rest with 0
	// In case of other numbers rather than 0, for instance if a number is even move it to the last position
	// but keep the order of the rest, the solution would be different
	start := time.Now()
	for index, elem := range elements {
		fmt.Printf("index %v", index)
		if elem == 0 {
			for i := index; i < len(elements)-1; i++ {
				elements[i] = elements[i+1]
			}
			elements[len(elements)-1] = 0
		}
	}
	fmt.Printf("Elements %v", elements)
	elapsed := time.Since(start)
	fmt.Println("Result:", elements)
	fmt.Printf("Execution time: %v\n", elapsed)

	elements = []int{6, 1, 0, 3, 0, 1, 0, 12}

	start = time.Now()

	pos := 0
	for _, elem := range elements {
		if elem != 0 {
			elements[pos] = elem
			pos++
		}
	}
	for i := pos; i < len(elements); i++ {
		elements[i] = 0
	}
	fmt.Printf("Elements %v", elements)
	elapsed = time.Since(start)
	fmt.Println("Result:", elements)
	fmt.Printf("Execution time: %v\n", elapsed)
}
