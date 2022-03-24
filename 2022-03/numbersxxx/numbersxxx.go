package main

import (
	"fmt"
)

func upground(n int) int {
	k := 9
	sum := 0
	bit := 0
	for k >= 0 && sum < n {
		sum += k
		k--
		bit++
	}
	if k < 0 && sum < n {
		return -1
	}
	return bit
}

func sum(nums []int, n int) (int, int) {
	if len(nums) < n {
		return -1, -1
	}
	d := 0
	for i := 0; i < n; i++ {
		d += nums[i]
	}
	u := 0
	for i := len(nums) - 1; i > len(nums)-n-1; i-- {
		u += nums[i]
	}
	return d, u
}

func find(target int, nums []int, n int) []int {
	if target == 0 {
		return []int{}
	}
	if target != 0 && len(nums) < n {
		return nil
	}
	down, up := sum(nums, n)
	if !(target >= down && target <= up) {
		return nil
	}
	for i := range nums {
		r := find(target-nums[i], nums[i+1:], n-1)
		if r != nil {
			return append([]int{nums[i]}, r...)
		}
	}
	return nil
}

func solve(n int) int {
	bits := upground(n)
	//fmt.Println(bits)
	if bits == -1 {
		return -1
	}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := find(n, nums, bits)
	if r == nil {
		return -1
	}
	ans := 0
	for i := range r {
		ans = ans*10 + r[i]
	}
	return ans
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	n = solve(n)
	fmt.Printf("%d\n", n)
}
