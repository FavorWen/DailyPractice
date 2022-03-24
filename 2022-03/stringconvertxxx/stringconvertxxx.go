package main

import (
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func input() (n int, a1 []int, a2 []int) {
	var s string
	fmt.Scanf("%d", &n)
	a1 = make([]int, n)
	a2 = make([]int, n)
	fmt.Scanf("%s", &s)
	for i := 0; i < n; i++ {
		a1[i] = int(s[i]) - 97
	}
	fmt.Scanf("%s", &s)
	for i := 0; i < n; i++ {
		a2[i] = int(s[i]) - 97
	}
	return n, a1, a2
}

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

func main() {
	n, a1, a2 := input()
	quickSort(a1)
	quickSort(a2)
	//fmt.Println(a1, a2)
	cost := 0
	for i := 0; i < n; i++ {
		if a1[i] != a2[i] {
			cost += abs(a2[i] - a1[i])
		}
	}
	fmt.Println(cost)
}
