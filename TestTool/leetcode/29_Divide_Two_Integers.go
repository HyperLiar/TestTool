package main

import "math"

/*
Given two integers dividend and divisor,
divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
 */

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if divisor == math.MinInt32 {
		if dividend == divisor {
			return 1
		} else {
			return 0
		}
	}

	minus := false
	if divisor < 0 {
		divisor = -divisor
		minus = !minus
	}

	if dividend < 0 {
		dividend = -dividend
		minus = !minus
	}

	count := 0
	rank, tmp := 1, divisor
	for dividend >= divisor && rank > 0 {
		for dividend >= tmp {
			dividend -= tmp
			count += rank
			if tmp < math.MaxInt32 {
				tmp <<= 1
				rank <<= 1
			}
		}
		rank, tmp = rank>>1, tmp>>1
	}

	if minus {
		return -count
	} else {
		return count
	}
}
