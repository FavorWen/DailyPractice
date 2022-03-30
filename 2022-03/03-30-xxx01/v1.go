package main

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	lo := 0
	hi := len(nums) - 1
	r := nums[lo]
	for lo < hi {
		for lo < hi && nums[hi] >= r {
			hi--
		}
		nums[lo] = nums[hi]
		for lo < hi && nums[lo] <= r {
			lo++
		}
		nums[hi] = nums[lo]
	}
	nums[lo] = r
	quickSort(nums[:lo])
	quickSort(nums[lo+1:])
}

func threeSumV1(nums []int) [][]int {
	quickSort(nums)
	r := make([][]int, 0)
	for first := 0; first < len(nums)-2; first++ {
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}
		if nums[first] > 0 {
			return r
		}
		third := len(nums) - 1
		for second := first + 1; second < len(nums)-1 && second < third; {
			if second != first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				r = append(r, []int{nums[first], nums[second], nums[third]})
				for second < third && nums[second] == nums[second+1] {
					second++
				}
				for second < third && nums[third] == nums[third-1] {
					third--
				}
				second++
				third--
			} else if nums[first]+nums[second]+nums[third] < 0 {
				second++
			} else {
				third--
			}
		}
	}
	return r
}
