package main

import (
	"fmt"
	"time"
)

func main() {
	elements := []string{"a", "b", "c", "b", "d", "a"}
	l := 0
	total := 1
	leftp := 0
	start := time.Now()
	for r := 0; r <= len(elements)-1; r++ {
		l = r - 1
		fmt.Printf("Max substring: %d\n", total)
		total = 1
		for l >= leftp {
			if elements[r] != elements[l] {
				l--
				total++
			} else {
				leftp = r - 1
				break
			}
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
	fmt.Printf("Max substring: %d\n", total)

	//subsMap := make(map[string]bool, len(elements))

	start = time.Now()
	//...
	elapsed = time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)

}
