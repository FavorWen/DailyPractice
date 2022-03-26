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

func mixSearch(nums []int, target int, lo int, hi int) int {
	if len(nums) == 0 {
		return -1
	}
	if lo > hi {
		return -1
	}
	if lo == hi {
		if nums[lo] == target {
			return lo
		} else {
			return -1
		}
	}
	var idx int
	mid := (lo + hi) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[lo] < nums[mid] {
		idx = quickSearch(nums[lo:mid], target)
		if idx != -1 {
			return idx + lo
		}
	} else if nums[lo] > nums[mid] {
		idx = mixSearch(nums, target, lo, mid-1)
		if idx != -1 {
			return idx
		}
	} else if nums[lo] == nums[mid] {
		if nums[lo] == target {
			return lo
		}
	}
	if nums[mid] < nums[hi] {
		idx = quickSearch(nums[mid:hi+1], target)
		if idx != -1 {
			return idx + mid
		}
	} else if nums[mid] > nums[hi] {
		idx = mixSearch(nums, target, mid+1, hi)
		if idx != -1 {
			return idx
		}
	} else if nums[mid] == nums[hi] {
		if nums[mid] == target {
			return mid
		}
	}
	return -1
}

func search(nums []int, target int) int {
	return mixSearch(nums, target, 0, len(nums)-1)
}
