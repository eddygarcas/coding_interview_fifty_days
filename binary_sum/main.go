package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []int32{1, 0, 1, 1}
	b := []int32{1, 1, 0, 1}
	result := addBinary(a, b)
	fmt.Printf("Addition result: %s\n", result)
}

func addBinary(a []int32, b []int32) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	result := ""
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
		result = strconv.Itoa(sum%2) + result
		carry = sum / 2

	}
	return result
}
