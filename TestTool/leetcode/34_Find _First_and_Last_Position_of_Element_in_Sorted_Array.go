package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	l := len(nums)
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	lt, gt := 0, l-1

	for lt <= gt {
		mid := int(uint(lt+gt)>>1)

		if nums[mid] == target {
			res[0] = searchLeftSide(nums, target, lt, mid)
			res[1] = searchRightSide(nums, target, mid, gt)
			break
		} else if nums[mid] > target {
			gt = mid -1
		} else {
			lt = mid + 1
		}
	}

	return res
}

func searchLeftSide(nums []int, target, lt, gt int) int {
	res := gt

	for lt <= gt {
		mid := int(uint(lt+gt)>>1)
		if nums[mid] < target {
			if mid+1 < len(nums) && nums[mid+1] == target {
				return mid+1
			}
			lt = mid+1
		} else {
			if target == nums[mid] {
				res = mid
			}
			gt = mid-1
		}
	}

	return res
}

func searchRightSide(nums []int, target, lt, gt int) int {
	res := lt

	for lt <= gt {
		mid := int(uint(lt+gt)>>1)
		if nums[mid] > target {
			if mid-1 >= 0 && nums[mid-1] == target {
				return mid-1
			}
			gt = mid -1
		} else {
			if nums[mid] == target {
				res = mid
			}
			lt = mid + 1
		}
	}

	return res
}