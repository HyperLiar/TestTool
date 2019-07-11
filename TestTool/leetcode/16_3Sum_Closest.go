package main

import (
	"math"
	"sort"
)

/**
Given an array nums of n integers and an integer target,
find three integers in nums such that the sum is closest to target.
Return the sum of the three integers.
You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */

func threeSumClosest(nums []int, target int) int {
	var res = nums[0]+nums[1]+nums[2]
	sort.Ints(nums)

	for i := 0; i < len(nums);i++ {
		lo, hi := i+1, len(nums)-1

		for lo < hi {
			cur := nums[i] + nums[lo] + nums[hi]
			if cur == target {
				return target
			} else if cur < target {
				lo++
			} else {
				hi--
			}

			d1 := math.Abs(float64(cur-target))
			d2 := math.Abs(float64(res-target))

			if d1 < d2 {
				res = cur
			}
		}
	}

	return res
}