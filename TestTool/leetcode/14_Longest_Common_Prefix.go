package main

/**
If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	var (
		i, j int
		str string
	)
	for i = 1; i < len(strs); i++ {
		str = strs[i]
		for j = 0; j < len(prefix) && j < len(str) && prefix[j] == str[j]; j++ {
		}
		prefix = prefix[:j]
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func main() {
	strs := []string{"aa", "a"}
	println(longestCommonPrefix(strs))
}
