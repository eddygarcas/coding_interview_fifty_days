package main

import "fmt"

func main() {
	keypad := make(map[int][]byte)
	keypad[2] = []byte{'a', 'b', 'c'}
	keypad[3] = []byte{'d', 'e', 'f'}
	keypad[4] = []byte{'g', 'h', 'i'}
	keypad[5] = []byte{'j', 'k', 'l'}
	keypad[6] = []byte{'m', 'n', 'o'}
	keypad[7] = []byte{'p', 'q', 'r', 's'}
	keypad[8] = []byte{'t', 'u', 'v'}
	keypad[9] = []byte{'w', 'x', 'y', 'z'}

	result := phoneCombination(keypad, func() []int { return []int{2, 3} })
	fmt.Printf("\n%c", result)
}

func phoneCombination(keypad map[int][]byte, input func() []int) []string {
	fmt.Printf("Keypad: %c\n", keypad)
	fmt.Printf("Input: %v\n", input())

	ans := []string{}

	digits := input()

	letters := keypad[digits[0]]

	for _, letter := range letters {
		for i := 1; i < len(digits); i++ {
			for _, other := range keypad[digits[i]] {
				ans = append(ans, string([]byte{letter, other}))
			}
		}
	}

	return ans
}
