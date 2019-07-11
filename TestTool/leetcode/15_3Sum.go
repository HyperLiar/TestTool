package main

import (
	"fmt"
	"sort"
)

/**
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

 */
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	var sum, lo, hi int

	sort.Ints(nums)

	set := make([][]int, 0)
	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lo, hi = i+1,len(nums)-1

		for lo < hi {
			sum = nums[lo] + nums[hi] + nums[i]

			if sum == 0 {
				set = append(set, []int{nums[i], nums[lo], nums[hi]})
				lo++
				hi--

				for lo <hi && nums[lo] == nums[lo-1] {
					lo++
				}

				for lo < hi && nums[hi] == nums[hi+1] {
					hi--
				}
			} else if sum >0 {
				hi--
			} else {
				lo++
			}
		}
	}

	return set
}

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
}