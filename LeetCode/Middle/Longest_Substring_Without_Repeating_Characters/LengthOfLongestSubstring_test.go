package lengthOfLongestSubstring

import (
	"testing"
)

type testingStruct struct {
	Data   string
	Result int
}

func Test_LengthOfLongestSubstring(t *testing.T) {
	testStucts := []testingStruct{
		{Data: "pwwkew", Result: 3},
		{Data: "ckilbkd", Result: 5},
		{Data: "asljlj", Result: 4},
		{Data: "abcabcbb", Result: 3},
		{Data: "abcb", Result: 3},
		{Data: "dvdf", Result: 3},
		{Data: "wobgrovw", Result: 6},
	}

	for _, tData := range testStucts {
		if LengthOfLongestSubstring(tData.Data) != tData.Result {
			t.Errorf("String %s. The answer is: %d. Output: %d",
				tData.Data,
				tData.Result,
				LengthOfLongestSubstring(tData.Data))
		}
	}
}
