package main

import "strings"

func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("bbbbb"))
	println(lengthOfLongestSubstring("pwwkew"))
	println(lengthOfLongestSubstring(" "))
	println(lengthOfLongestSubstring("dvdf"))
	println(lengthOfLongestSubstring("asjrgapa"))
	println(lengthOfLongestSubstring("pwwkew"))

}

func lengthOfLongestSubstring(s string) int {
	var longestSubstr = 0
	var substr = ""
	for i := range s {
		for _, char := range s[i:] {
			if !strings.Contains(substr, string(char)) {
				substr += string(char)
			} else {
				break
			}
			if longestSubstr < len(substr) {
				longestSubstr = len(substr)
			}
		}
		substr = ""
	}
	return longestSubstr
}
