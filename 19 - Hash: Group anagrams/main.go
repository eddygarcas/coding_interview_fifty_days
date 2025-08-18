package main

import (
	"fmt"
	"sort"
	"time"
)

// main is the entry point of the program.
// It demonstrates two methods to group anagrams from a list of strings and benchmarks their execution time.
func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := make([][]string, len(input))

	start := time.Now()
	output = groupAnagrams(input)
	elapsed := time.Since(start)
	fmt.Printf("Output: %v\n", output)
	fmt.Printf("Group Anagrams took %s\n", elapsed)

	start = time.Now()
	output = groupAnagrams2(input)
	elapsed = time.Since(start)
	fmt.Printf("Output: %v\n", output)
	fmt.Printf("Group Anagrams took %s\n", elapsed)
}

// groupAnagrams groups a slice of strings into anagrams using character comparison.
// Time Complexity: O(n^2 * k), where n is the number of strings and k is the average string length.
// Space Complexity: O(nk) for the auxiliary map and result slices.
// This implementation uses a nested loop and a helper function to compare character counts.
func groupAnagrams(anagrams []string) [][]string {
	// aux tracks which strings have been grouped
	aux := make(map[string]int)
	// result stores grouped anagrams
	result := make([][]string, len(anagrams))
	l := 0
	for l < len(anagrams) {
		for _, anagram := range anagrams[l:] {
			if aux[anagram] == 0 {
				if compareChars(anagrams[l], anagram) {
					result[l] = append(result[l], anagram)
					aux[anagram] = l + 1
				}
			}
		}
		l++
	}
	return result
}

// compareChars checks if two strings are anagrams by comparing their character counts.
// Time Complexity: O(k), where k is the length of the strings.
// Space Complexity: O(k) for the character count map.
// Returns true if both strings contain the same characters with the same frequencies.
func compareChars(s, t string) bool {
	// count stores the frequency of each character in string s
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, c := range t {
		if (count[c]) == 0 {
			return false
		}
		count[c]--
	}
	// Check if all character counts are zero
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

// groupAnagrams2 groups a slice of strings into anagrams by sorting their characters.
// For each string of length k, sort.Slice() takes O(k log k) time, so for n strings:
// Time Complexity: O(n * k log k), where n is the number of strings and k is the average string length.
// Space Complexity: O(nk) for storing the map and result slices.
// It uses a map where the key is the sorted string and the value is a list of anagrams.
// Returns a slice of string slices, where each inner slice contains words that are anagrams of each other.
func groupAnagrams2(anagrams []string) [][]string {
	// resultMap stores the sorted strings as keys and their corresponding anagrams as values
	resultMap := make(map[string][]string)
	for _, anagram := range anagrams {
		array := []byte(anagram)
		sort.Slice(array, func(i, j int) bool {
			return array[i] < array[j]
		})
		sortedS := string(array)
		if _, ok := resultMap[sortedS]; !ok {
			resultMap[sortedS] = make([]string, 0)
		}
		resultMap[sortedS] = append(resultMap[sortedS], anagram)
	}
	// Initialize result with the correct capacity
	result := make([][]string, 0, len(resultMap))
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}
