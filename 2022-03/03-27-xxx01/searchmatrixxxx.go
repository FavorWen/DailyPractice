package main

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	lo := 0
	hi := row - 1
	for lo < hi-1 {
		if matrix[lo][0] > target || matrix[hi][col-1] < target {
			return false
		}
		mid := (lo + hi) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	if lo == hi-1 {
		if matrix[hi][0] > target {
			row = lo
		} else {
			row = hi
		}
	} else {
		row = lo
	}
	lo = 0
	hi = col - 1
	//fmt.Println(row)
	for lo < hi {
		mid := (lo + hi) / 2
		if matrix[row][mid] < target {
			lo = mid + 1
		} else if matrix[row][mid] > target {
			hi = mid - 1
		} else {
			return true
		}
	}
	if lo == hi && matrix[row][lo] == target {
		return true
	}
	return false
}

func v2(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	lo := 0
	hi := row - 1
	for lo <= hi {
		mid := (lo + hi) / 2
		if target >= matrix[mid][0] && target <= matrix[mid][col-1] {
			row = mid
			break
		} else if target < matrix[mid][0] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if lo > hi {
		return false
	}
	lo = 0
	hi = col - 1
	for lo <= hi {
		mid := (lo + hi) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func main() {
	nums := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	searchMatrix(nums, 30)

}
