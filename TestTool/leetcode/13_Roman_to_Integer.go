package main

/**
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
 */
func romanToInt(s string) int {
	romanValue := map[int32]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	num := 0
	for key, val := range s {
		if key < len(s)-1 {
			if s[key] == 'I' && (s[key+1] == 'V' || s[key+1] == 'X') {
				num += -romanValue[val]
			} else if s[key] == 'X' && (s[key+1] == 'L' || s[key+1] == 'C') {
				num += -romanValue[val]
			} else if s[key] == 'C' && (s[key+1] == 'D' || s[key+1] == 'M') {
				num += -romanValue[val]
			} else {
				num += romanValue[val]
			}
		} else {
			num += romanValue[val]
		}
	}

	return num
}
