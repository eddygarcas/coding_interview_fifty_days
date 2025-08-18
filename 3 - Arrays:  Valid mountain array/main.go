package main

// This program checks a given array to determine if it forms a valid mountain array.
// A mountain array increases to a single peak and then decreases, with at least three elements.

import "fmt"

func main() {
	// mountain represents the input array to check
	var mountain = []int{0, 4, 5, 6, 3, 2, 1, 1}
	if len(mountain) < 3 {
		fmt.Println("mountain less than 3")
	}
	// first will check the mountain up
	// Check if there are more element
	//if they do, check if it's going down
	// Find the index and value of the peak (maximum element)
	maxIndex := 0
	maxValue := 0
	for index, elem := range mountain {
		if elem > maxValue {
			maxValue = elem
			maxIndex = index
		}
	}
	fmt.Println(maxIndex)
	// Check if the peak is not at the start or end
	if (maxIndex == len(mountain)-1) || (maxIndex == 0) {
		fmt.Println("array has no elements to cover the array after the decrement")
		return
	}
	// Validate that all elements except the peak are less than the peak
	valid := false
	for i := 0; i < len(mountain); i++ {
		if i == maxIndex {
			continue
		}
		if mountain[i] < maxValue {
			valid = true
		} else {
			valid = false
			break
		}
	}
	fmt.Println(valid)
}
