package main

import "fmt"

func main() {
	var mountain = []int{0, 4, 5, 6, 3, 2, 1, 1}
	if len(mountain) < 3 {
		fmt.Println("mountain less than 3")
	}
	// first will check the mountain up
	// Check if there are more element
	//if they do, check if it's going down
	maxIndex := 0
	maxValue := 0
	for index, elem := range mountain {
		if elem > maxValue {
			maxValue = elem
			maxIndex = index
		}
	}
	fmt.Println(maxIndex)
	if (maxIndex == len(mountain)-1) || (maxIndex == 0) {
		fmt.Println("array has no elements to cover the array after the decrement")
		return
	}
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
