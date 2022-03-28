package main

import "fmt"

func isPeak(nums []int, idx int) bool {
	if len(nums) == 1 {
		return true
	}
	if idx == 0 {
		if nums[0] > nums[1] {
			return true
		} else {
			return false
		}
	}
	if idx == len(nums)-1 {
		if nums[idx] > nums[idx-1] {
			return true
		} else {
			return false
		}
	}
	if nums[idx-1] < nums[idx] && nums[idx+1] < nums[idx] {
		return true
	}
	return false
}

func isUp(nums []int, idx int) bool {
	if idx == 0 {
		return true
	}
	if idx == len(nums)-1 {
		return false
	}
	return nums[idx-1] < nums[idx]
}

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}
	if len(nums) == 2 {
		return -1
	}
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := (lo + hi) / 2
		if isPeak(nums, mid) {
			return mid
		}
		if isUp(nums, mid) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if isPeak(nums, lo) {
		return lo
	}
	return -1
}

func main() {
	n := []int{3, 4, 3, 2, 1}
	fmt.Println(findPeakElement(n))
}
