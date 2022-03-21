package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	H     int
	W     int
	Area  int
	ratio float64
}

func less(m1, m2 *Matrix) bool {
	if m1.Area < m2.Area {
		return true
	} else if m1.Area > m2.Area {
		return false
	} else if m1.ratio > m2.ratio {
		return true
	} else if m1.ratio < m2.ratio {
		return false
	} else if m1.W < m2.W {
		return true
	} else if m1.W > m2.W {
		return false
	} else {
		return false
	}
}

func swap(ms []Matrix, i, j int) {
	m := Matrix{}
	m = ms[i]
	ms[i] = ms[j]
	ms[j] = m
}

func quickSort1(ms []Matrix) {
	if len(ms) < 2 {
		return
	}
	r := ms[0]
	st := ms[0]
	i := 0
	j := len(ms) - 1
	for i < j {
		// see equal as less
		for i < j && less(&st, &ms[j]) {
			j--
		}
		ms[i] = ms[j]
		// see equal as less
		for i < j && !less(&st, &ms[i]) {
			i++
		}
		ms[j] = ms[i]
	}
	ms[i] = r
	quickSort1(ms[0:i])
	quickSort1(ms[i+1:])
}

func quickSort2(ms []Matrix) {
	if len(ms) < 2 {
		return
	}
	st := ms[0]
	i := 0
	j := len(ms) - 1
	for i < j {
		for i < j && less(&st, &ms[j]) {
			j--
		}
		swap(ms, i, j)
		for i < j && !less(&st, &ms[i]) {
			i++
		}
		swap(ms, i, j)
	}
	quickSort2(ms[0:i])
	quickSort2(ms[i+1:])
}

func main() {

	// input
	var n int
	matrixs := make([]Matrix, 0)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var h, w int
		fmt.Scanf("%d%d", &w, &h)
		m := Matrix{
			H:     h,
			W:     w,
			Area:  h * w,
			ratio: math.Min(float64(h)/float64(w), float64(w)/float64(h)),
		}
		matrixs = append(matrixs, m)
	}

	// Sort
	// quickSort1 is better
	//quickSort1(matrixs)
	quickSort2(matrixs)

	// output
	for _, m := range matrixs {
		fmt.Printf("%d %d ", m.W, m.H)
	}
	fmt.Printf("\n")
}
