package main

import "math"

// O(n^2)
func maxArea(height []int) int {
	var max = 0

	for i := 0;i < len(height)-1; i++ {
		for j := i+1; j < len(height); j++ {
			now := minInt(height[j], height[i]) * (j - i)

			if now > max {
				max = now
			}
		}
	}

	return max
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	}

	max := 0

	for i,j := 0,l-1;i<j; {
		h := int(math.Min(float64(height[i]), float64(height[j])))

		now := h * (j - i)
		if now > max {
			max = now
		}

		if height[j] > height[i] {
			i++
		} else {
			j--
		}
	}

	return max
}