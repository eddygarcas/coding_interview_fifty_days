package main

import (
	"fmt"
	"time"
)

// position represents a point in 2D space with x and y coordinates
type position struct {
	x, y int
}

func main() {
	start := time.Now()
	robot := position{0, 0}
	sequence := []byte{'U', 'R', 'L', 'D', 'U'}

	// Calculate final robot position after sequence of moves
	robot.finalPosition(sequence)
	fmt.Printf("Final position: %v\n", robot)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}

// finalPosition calculates the robot's final position after executing a sequence of moves
// The moves are:
// U = Move up (y++)
// D = Move down (y--)
// R = Move right (x++)
// L = Move left (x--)
// Time complexity: O(n) where n is length of sequence
// Space complexity: O(1) as map stores max 4 directions
func (p *position) finalPosition(seq []byte) {
	pos := make(map[byte]int, 4)
	for _, val := range seq {
		pos[val]++
	}
	p.x = pos['R'] - pos['L']
	p.y = pos['U'] - pos['D']
}
