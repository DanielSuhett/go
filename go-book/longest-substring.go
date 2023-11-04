package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var bytes = make(map[byte]int)
	var value int

	for _, v := range []byte(s) {
		if bytes[v] == 0 {
			value++
		}
		bytes[v] = bytes[v] + 1
	}

	fmt.Println(value)

	return value
}

func init() {
	lengthOfLongestSubstring("abcbb")
	lengthOfLongestSubstring("bbbbbbbb")
	lengthOfLongestSubstring("pwwkew")

}
