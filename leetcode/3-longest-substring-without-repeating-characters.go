package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)

	maxLen := 0
	start := -1 //
	for end, r := range s {
		if last, ok := m[r]; ok && start < last {
			start = last
		}
		if end-start > maxLen {
			maxLen = end - start
		}
		m[r] = end
	}
	return maxLen

}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
