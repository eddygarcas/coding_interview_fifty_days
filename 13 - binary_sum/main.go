package main

// This program adds two binary numbers represented as slices of int32 and returns their sum as a string.
// It simulates binary addition, handling carry-over for each digit from right to left.

import (
	"fmt"
	"strconv"
)

func main() {
	// Example binary numbers: a = 1011, b = 1101
	a := []int32{1, 0, 1, 1}
	b := []int32{1, 1, 0, 1}
	result := addBinary(a, b)
	fmt.Printf("Addition result: %s\n", result)
}

// addBinary adds two binary numbers represented as slices and returns the sum as a string.
func addBinary(a []int32, b []int32) string {
	i := len(a) - 1 // Pointer for a
	j := len(b) - 1 // Pointer for b
	carry := 0      // Carry for addition
	result := ""    // Resulting binary string
	// Loop until both pointers are exhausted and no carry remains
	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i])
			i--
		}
		if j >= 0 {
			sum += int(b[j])
			j--
		}
		// Prepend the current binary digit to the result
		result = strconv.Itoa(sum%2) + result
		// Update carry for the next iteration
		carry = sum / 2
	}
	return result
}
