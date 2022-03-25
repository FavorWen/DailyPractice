package main

import (
	"fmt"
	"strconv"
)

func input() (int, []int) {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
		nums[i] = nums[i] % m
	}
	return m, nums
}

func mk(x, y int) string {
	return strconv.Itoa(x) + "-" + strconv.Itoa(y)
}

// S(i,j) = Ai + S(i+1, j-1) + Aj

func S(x int, y int, nums []int, pre map[string]int, m int) int {
	if x == y {
		return nums[x]
	} else {
		if x == y-1 {
			return (nums[x] + nums[y]) % m
		} else {
			sij := pre[mk(x+1, y-1)]
			return (nums[x] + sij + nums[y]) % m
		}
	}
}

func solve(m int, nums []int) int {
	repre := make(map[string]int)
	pre := make(map[string]int)
	next := make(map[string]int)
	r := 0
	for length := 1; length <= len(nums); length++ {
		for i := 0; i+length <= len(nums); i++ {
			sum := S(i, i+length-1, nums, repre, m)
			next[mk(i, i+length-1)] = sum
			if sum == 0 {
				r++
			}
		}
		repre = pre
		pre = next
		next = make(map[string]int)
	}
	return r
}

func main() {
	m, nums := input()
	n := solveV1(m, nums)
	fmt.Println(n)
}
