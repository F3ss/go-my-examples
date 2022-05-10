package lengthOfLongestSubstring

func LengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	chars := make(map[rune]int)
	result2 := 1
	result1 := 0

	for index, char := range s {
		if _, ok := chars[char]; !ok {
			result1++
			chars[char] = index
		} else {
			if result1 > result2 {
				result2 = result1
			}
			result1 = index - chars[char]
			for key, val := range chars {
				if val < chars[char] {
					delete(chars, key)
				}
			}
			chars[char] = index
		}
	}

	if result1 < result2 {
		result1 = result2
	}

	return result1
}
