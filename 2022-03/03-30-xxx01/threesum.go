package main

import (
	"fmt"
	"strconv"
)

func mkstr(i, j, k int) string {
	if i > j {
		i, j = j, i
	}
	if j > k {
		j, k = k, j
	}
	if i > j {
		i, j = j, i
	}
	return strconv.Itoa(i) + "*" + strconv.Itoa(j) + "*" + strconv.Itoa(k)
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	record := make(map[string]struct{})
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					if _, ok := record[mkstr(nums[i], nums[j], nums[k])]; !ok {
						res = append(res, []int{nums[i], nums[j], nums[k]})
						record[mkstr(nums[i], nums[j], nums[k])] = struct{}{}
					}
				}
			}
		}
	}
	return res
}

func main() {
	n := []int{-1, 0, 1, 2, -1, -4}
	// k := threeSum(n)
	// for _, i := range k {
	// 	fmt.Println(i)
	// }
	fmt.Println(threeSumV1(n))
}
