package main

import (
	"fmt"
	"sort"
)

func main() {
	// Remember only TWO person can fit on a boat regardless it's weight
	people := []int{3, 3, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1}
	limit := 3 // 7
	boats := 0
	sort.Ints(people)
	fmt.Println("Ordered people", people)
	//Ordered people [3 3 2 2 2 2 2 1 1 1 1 1 1 1 1]
	r := len(people) - 1
	l := 0
	for r >= l {
		if people[r]+people[l] <= limit {
			r--
			l++
		} else {
			r--
		}
		boats++
	}
	fmt.Printf("Total number of boats %v", boats)
}
