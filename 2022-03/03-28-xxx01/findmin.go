package main

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return min(nums[0], nums[1])
	}
	lo := 0
	hi := len(nums) - 1
	if nums[lo] < nums[hi] {
		return nums[lo]
	}
	mid := (lo + hi) / 2
	return min(findMin(nums[lo:mid]), findMin(nums[mid:]))
}
