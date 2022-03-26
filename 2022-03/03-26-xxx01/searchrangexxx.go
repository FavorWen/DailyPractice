package main

func quickSearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := (lo + hi) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if lo != hi {
		return -1
	}
	if nums[lo] == target {
		return lo
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	idx := quickSearch(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}
	var lo, hi int
	for lo = idx; lo >= 0 && nums[lo] == target; {
		lo--
	}
	for hi = idx; hi < len(nums) && nums[hi] == target; {
		hi++
	}
	return []int{lo + 1, hi - 1}
}
