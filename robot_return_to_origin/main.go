package main

import (
	"fmt"
	"time"
)

type position struct {
	x, y int
}

func main() {

	start := time.Now()
	robot := position{0, 0}
	sequence := []byte{'U', 'R', 'L', 'D', 'U'}
	/**
	U = y++
	D = y--
	R = x++
	L = x--
	You can implement that just return true if the robot goes back to 0,0
	also you can just use a switch to increment or decrement the values to achieve O(N) on
	time complexity, and O(1) in space complexity.
	Anyway here I've decides to use map as well as position struct to return the position as well as
	use a common representation for 2D spaces
	*/
	robot.finalPosition(sequence)
	fmt.Printf("Final position: %v\n", robot)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\v", elapsed)
}

func (p *position) finalPosition(seq []byte) {
	//create a hashmap for every direction
	pos := make(map[byte]int, 4)
	for _, val := range seq {
		pos[val] += 1
	}
	p.x = pos['R'] - pos['L']
	p.y = pos['U'] - pos['D']
}
