package main

// 解析
func solveV1(m int, nums []int) int {
	r := 0
	sum := 0
	record := make(map[int]int)
	for i := 1; i < m; i++ {
		record[i] = 0
	}
	record[0] = 1
	for i := 0; i < len(nums); i++ {
		sum = (sum + nums[i]) % m
		r += record[sum]
		record[sum]++
	}
	return r
}
