package main

// This program solves the Group Anagrams problem using hash tables.
// It takes a list of strings and groups them based on whether they are anagrams of each other.
// Two strings are anagrams if they contain the same characters with the same frequencies.
//
// Example:
// Input: ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// The program implements two different approaches:
// 1. Character counting using a hash map (groupAnagrams)
// 2. Sorting characters and using sorted string as key (groupAnagrams2)

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
// For each string, it compares with others to find anagrams by counting characters.
//
// Time Complexity: O(n^2 * k), where n is the number of strings and k is the average string length
// Space Complexity: O(nk) for storing the auxiliary map and result slices
func groupAnagrams(anagrams []string) [][]string {
	aux := make(map[string]int)
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
// It uses a hash map to track character frequencies.
//
// Time Complexity: O(k), where k is the length of the strings
// Space Complexity: O(1) since character set is fixed
func compareChars(s, t string) bool {
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
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

// groupAnagrams2 groups anagrams by sorting each string's characters.
// Uses sorted strings as keys in a hash map to group anagrams efficiently.
//
// Time Complexity: O(n * k log k), where n is number of strings, k is average string length
// Space Complexity: O(nk) for storing the map and result slices
func groupAnagrams2(anagrams []string) [][]string {
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
	result := make([][]string, 0, len(resultMap))
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}
