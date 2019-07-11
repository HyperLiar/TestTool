package main

import "fmt"

func removeDuplicates(nums []int) int {
	start := false
	var before int
	keyNum,idx := 0,1

	for _, v := range nums {
		if !start {
			before = v
			start = true
			keyNum++
		} else if before != v {
			before = v
			nums[idx] = v
			idx++
			keyNum++
		}
	}

	return keyNum
}

func removeDuplicates(nums []int) int {
	for i, j := 0, 1; i < j && j < len(nums); {
		if nums[j] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			j--
			i--
		}
		j++
		i++
	}
	return len(nums)
}

func main() {
	nums := []int{1,1,2}
	fmt.Println(removeDuplicates(nums))
}