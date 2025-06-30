package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	start := time.Now()
	e := []byte{'a', 'c', 'c', 'd', 'e'}
	subsMap := make(map[byte]int, len(e))

	ans := 0
	left := 0
	right := 0
	for i := 0; i < len(e); i++ {
		if val, ok := subsMap[e[i]]; ok {
			left = int(math.Max(float64(left), float64(val+1)))
		}
		right++
		subsMap[e[i]] = i
		ans = int(math.Max(float64(right-left), float64(ans)))
	}
	fmt.Printf("Max substring: %d\n", ans)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)

}
