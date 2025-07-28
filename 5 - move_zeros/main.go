package main

// This program demonstrates two methods to move all zeros in an array
// to the end while maintaining the order of non-zero elements.
// The first method shifts zeros one by one, which is less efficient.
// The second method uses a two-pass approach for better performance.

import (
	"fmt"
	"time"
)

func main() {
	// Initial array with zeros and non-zeros
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
		if elem == 0 {
			for i := index; i < len(elements)-1; i++ {
				elements[i] = elements[i+1]
			}
			elements[len(elements)-1] = 0
		}
	}
	fmt.Println("Elements: ", elements)
	elapsed := time.Since(start)
	fmt.Println("Result:", elements)
	fmt.Printf("Execution time: %v\n", elapsed)

	// Reset the array for the next method
	elements = []int{6, 1, 0, 3, 0, 1, 0, 12}

	// Method 2: Efficiently move non-zeros forward, then fill the rest with zeros
	start = time.Now()
	pos := 0 // Position to place the next non-zero element
	for _, elem := range elements {
		if elem != 0 {
			elements[pos] = elem
			pos++
		}
	}
	// Fill the remaining positions with zeros
	for i := pos; i < len(elements); i++ {
		elements[i] = 0
	}
	fmt.Println("Elements: ", elements)
	elapsed = time.Since(start)
	fmt.Println("Result:", elements)
	fmt.Printf("Execution time: %v\n", elapsed)
}
