package main

func removeElement(nums []int, val int) int {
	pre, idx := -1, 0

	for idx < len(nums) {
		if nums[idx] != val {
			pre++
			nums[pre], nums[idx] = nums[idx], nums[pre]
		}

		idx++
	}
	return pre+1
}

func main() {

}
