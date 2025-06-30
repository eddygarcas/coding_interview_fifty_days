package main

import "log"

func main() {
	var height = []int{1, 8, 6, 2, 5, 7}
	l := 0
	r := len(height) - 1
	log.Printf("height %v lengt %v", l, r)
	maxArea := 0
	for l < r {
		log.Printf("l %v r %v", l, r)
		width := r - l
		if height[l] < height[r] {
			maxArea = max(maxArea, width*height[l])
			l++
		} else {
			maxArea = max(maxArea, width*height[r])
			r--
		}
	}
	log.Printf("maxArea %v", maxArea)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
