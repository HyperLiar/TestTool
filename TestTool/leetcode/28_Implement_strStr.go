package main

// KMP
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	fail := failTable(needle)
	toReturn := -1
	i, j := 0, 0
	for i <= (len(haystack) - len(needle)) {
		for (j < len(needle)) && (needle[j] == haystack[i+j]) {
			j++
		}
		if j == 0 {
			i++
		} else {
			if j == len(needle) {
				return i
			}
			i = i + j - fail[j-1]
			j = fail[j-1]
		}
	}
	return toReturn
}

func failTable(p string) []int {
	l := len(p)
	t := make([]int, l)
	t[0] = 0
	i, j := 0, 1
	for j < l {
		if p[i] == p[j] {
			t[j] = i + 1
			i++
			j++
		} else {
			if i != 0 {
				i = t[i-1]
			} else {
				t[j] = i
				j++
			}
		}
	}
	return t
}
