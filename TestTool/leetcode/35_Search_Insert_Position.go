package main

import "fmt"

func searchInsert(nums []int, target int) int {
	lt, gt := 0, len(nums)-1

	for lt <= gt {
		mid := int(uint(lt+gt) >> 1)

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			gt = mid - 1
		} else {
			lt = mid + 1
		}
	}

	return lt
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
}
