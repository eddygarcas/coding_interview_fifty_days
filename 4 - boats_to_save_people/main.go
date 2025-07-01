package main

// This program solves the "Boats to Save People" problem.
// It determines the minimum number of boats required to carry everyone,
// given each boat can carry at most two people without exceeding a weight limit.

import (
	"fmt"
	"sort"
)

func main() {
	// people represents the weights of each person
	// Only two people can fit on a boat, regardless of their weight
	people := []int{3, 3, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1}
	limit := 3 // Maximum allowed weight per boat
	boats := 0 // Counter for number of boats used
	// Sort the array to efficiently pair the lightest and heaviest
	sort.Ints(people)
	fmt.Println("Ordered people", people)
	// Use two pointers to try to fit the lightest and heaviest together
	r := len(people) - 1 // Right pointer (heaviest)
	l := 0               // Left pointer (lightest)
	for r >= l {
		if people[r]+people[l] <= limit {
			// If the lightest and heaviest together fit, move both pointers
			r--
			l++
		} else {
			// Otherwise, the heaviest goes alone
			r--
		}
		boats++
	}
	fmt.Printf("Total number of boats %v", boats)
}
