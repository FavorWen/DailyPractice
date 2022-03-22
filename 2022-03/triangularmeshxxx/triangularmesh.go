package main

import (
	"fmt"
)

type Triangle struct {
	A, B, C  int
	Sequence bool //0 Clockwise, or counterClockwise
}

func input() []Triangle {
	var N, M int
	fmt.Scanf("%d%d", &N, &M)
	ts := make([]Triangle, M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d%d%d", &ts[i].A, &ts[i].B, &ts[i].C)
	}
	return ts
}

func insert(t *Triangle, record map[int]struct{}, seq bool) {
	if seq {
		key := t.A*1000 + t.B
		record[key] = struct{}{}
		key = t.B*1000 + t.C
		record[key] = struct{}{}
		key = t.C*1000 + t.A
		record[key] = struct{}{}
	} else {
		key := t.B*1000 + t.A
		record[key] = struct{}{}
		key = t.C*1000 + t.B
		record[key] = struct{}{}
		key = t.A*1000 + t.C
		record[key] = struct{}{}
	}
}

func find(t *Triangle, record map[int]struct{}) bool {
	key := t.A*1000 + t.B
	_, ok := record[key]
	if ok {
		t.Sequence = false
		insert(t, record, false)
		return true
	}
	key = t.B*1000 + t.C
	_, ok = record[key]
	if ok {
		t.Sequence = false
		insert(t, record, false)
		return true
	}
	key = t.C*1000 + t.A
	_, ok = record[key]
	if ok {
		t.Sequence = false
		insert(t, record, false)
		return true
	}

	key = t.B*1000 + t.A
	_, ok = record[key]
	if ok {
		t.Sequence = true
		insert(t, record, true)
		return true
	}
	key = t.C*1000 + t.B
	_, ok = record[key]
	if ok {
		t.Sequence = true
		insert(t, record, true)
		return true
	}
	key = t.A*1000 + t.C
	_, ok = record[key]
	if ok {
		t.Sequence = true
		insert(t, record, true)
		return true
	}
	return false
}

func solve(ts []Triangle) {
	undo := make([]int, len(ts)-1)
	for i := 1; i < len(ts); i++ {
		undo[i-1] = i
	}
	record := make(map[int]struct{})
	insert(&ts[0], record, true)
	ts[0].Sequence = true

	for idx := 0; len(undo) != 0; idx = (idx + 1) % len(undo) {
		i := undo[idx]
		if find(&ts[i], record) {
			undo = append(undo[:idx], undo[idx+1:]...)
		}
		if len(undo) == 0 {
			break
		}
	}

}

func minIndex(s []int) int {
	var min, minIndex int
	min = s[0]
	minIndex = 0
	if s[1] < min {
		min = s[1]
		minIndex = 1
	}
	if s[2] < min {
		minIndex = 2
	}
	return minIndex
}

func mod(x int) int {
	x = x % 3
	if x < 0 {
		x += 3
	}
	return x
}

func output(ts []Triangle) {
	for _, t := range ts {
		s := []int{t.A, t.B, t.C}
		midx := minIndex(s)
		if t.Sequence {
			fmt.Printf("%d %d %d\n", s[midx], s[mod(midx+1)], s[mod(midx+2)])
		} else {
			fmt.Printf("%d %d %d\n", s[midx], s[mod(midx-1)], s[mod(midx-2)])
		}
	}
}

func main() {
	ts := input()
	solve(ts)
	output(ts)
}
