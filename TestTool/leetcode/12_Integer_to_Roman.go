package main

import "strings"

/**
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
 */
func intToRoman(num int) string {
	bytes := []string{"I", "V", "X", "L", "C", "D", "M"}
	idx := 0
	res := ""
	for ; num > 0; num, idx = num/10, idx+1 {
		i := num % 10
		if i < 4 {
			res = strings.Repeat(bytes[2*idx], i) + res
		} else if i == 4 {
			res = bytes[2*idx] + bytes[2*idx+1] + res
		} else if i == 5 {
			res = bytes[2*idx+1] + res
		} else if i < 9 {
			res = bytes[2*idx+1] + strings.Repeat(bytes[2*idx], i-5) + res
		} else {
			res = bytes[2*idx] + bytes[2*idx+2] + res
		}
	}

	return res
}

func main() {
	tests := []int{3, 4, 9, 58, 1994}

	for _, val := range tests {
		println(intToRoman(val))
	}
}
