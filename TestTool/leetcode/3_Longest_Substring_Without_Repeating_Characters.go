package main

import "fmt"

func main() {
	//stringMap := []string{"abcabcbb", "bbbbb", "pwwkew", "aab"}
	stringMap := []string{" "}
	for _, s := range stringMap {
		fmt.Println(lengthOfLongestSubstring(s))
	}
}

func lengthOfLongestSubstring(s string) int {
	charMap := [128]int{}

	// 全部赋值 -1
	for key, _ := range charMap {
		charMap[key] = -1
	}

	maxLen, start := 0, 0

	for i, c := range s {
		// 出现过，至少是0，如果当前c已经出现过，对应的下标 >= start，设置start为这个下标+1
		if charMap[c] >= start {
			start = charMap[c]+1
		}

		// i赋值
		charMap[c] = i

		// i-start+1，为字符串的长度
		maxLen = maxInt(maxLen, i - start + 1)
	}

	return maxLen
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}