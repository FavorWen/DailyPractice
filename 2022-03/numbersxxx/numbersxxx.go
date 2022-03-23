package main

import (
	"fmt"
)

func upground(n int) (int, int, int) {
	r := 0
	k := 9
	sum := 0
	bit := 0
	for k >= 0 && sum < n {
		r = r*10 + k
		sum += k
		k--
		bit++
	}
	if k < 0 && sum < n {
		return -1, bit, 0
	}
	return r, bit, sum
}

func downground(up, bits, sum int) int {
	r := 0
	k := 1
	for ; bits > 0; bits-- {
		r = r*10 + k
		k++
	}
	return r
}

func solve(n int) int {
	up, bits, sum := upground(n)
	if sum == n {
		return up
	}
	down := downground(up, bits, sum)
	//fmt.Println(up, down, bits)
	for i := down; i <= up; i++ {
		sum := 0
		pool := make(map[int]int)
		flag := false
		k := i
		for ; k != 0; k = k / 10 {
			j := k % 10
			if _, ok := pool[j]; ok {
				flag = true
				break
			}
			sum += j
			if sum > n {
				flag = true
				break
			}
			pool[j] = 0
		}
		if flag {
			continue
		}
		if sum == n {
			return i
		}
	}
	return -1
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	n = solve(n)
	fmt.Printf("%d\n", n)
}
