package main

func main() {
	nums := []int{4,5,6,7,0,1,2}
	println(search(nums, 4))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	if r - l < 3 {
		return normalSearch(nums, target, l, r)
	}
	return specialSearch(nums, target, l, r)
}

func specialSearch(nums []int, target, l, r int) int {
	if r < l || r >= len(nums) || l < 0 {
		return -1
	}

	if r - l < 3 {
		return normalSearch(nums, target, l, r)
	}

	mid := (l + r) / 2

	if nums[mid] == target {
		return mid
	}
	if nums[l] < nums[mid] {
		if target <= nums[mid] && target >= nums[l] {
			return binarySearch(nums, target, l, mid-1)
		} else {
			return specialSearch(nums, target, mid+1, r)
		}
	} else {
		if target >= nums[mid] && target <= nums[r] {
			return binarySearch(nums, target, mid+1, r)
		} else {
			return specialSearch(nums, target, l, mid-1)
		}
	}
}

func binarySearch(nums []int, target, l, r int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func normalSearch(nums []int, target, l, r int) int {
	for i:=l; i <= r; i++ {
		if target == nums[i] {
			return i
		}
	}

	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}
	s, e := 0, len(nums)-1
	for s+1 < e{
		m := (s+e)/2
		if target == nums[m] {
			return m
		}
		if nums[s] < nums[m]{
			//left is sorted
			if target >=  nums[s] && target < nums[m]{
				e = m-1
			}else{
				s = m+1
			}
		}else {
			//right is sorted
			if target <= nums[e] && target > nums[m]{
				s = m+1
			}else{
				e = m-1
			}
		}
	}
	if nums[s] == target{
		return s
	}
	if nums[e] == target{
		return e
	}
	return -1
}