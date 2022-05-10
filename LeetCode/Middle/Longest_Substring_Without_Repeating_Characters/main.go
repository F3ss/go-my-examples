package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("ckilbkd"))
	fmt.Println(lengthOfLongestSubstring("asljlj"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // "abcb" //""
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	var r rune
	chars := make(map[rune]int)
	result2 := 1
	result1 := 0

	for index, char := range s {
		if char > 64 && char < 91 || char > 96 && char < 123 {
			if _, ok := chars[char]; !ok {
				result1++
				chars[char] = index
			} else if ok && char != r && len(s) > index+1 {
				if result1 < index-chars[char] {
					result1 = index - chars[char]
				}
				chars[char] = index

				if result1 > result2 {
					result2 = result1
				}
			} else {
				for key := range chars {
					delete(chars, key)
				}
				chars[char] = index

				if result1 > result2 {
					result2 = result1
				}
				result1 = 1
			}
		}
		r = char
	}

	if result1 < result2 {
		result1 = result2
	}

	return result1
}
