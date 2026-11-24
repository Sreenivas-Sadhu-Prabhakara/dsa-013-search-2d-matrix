package main

import "fmt"

// SearchMatrix reports whether target exists in a matrix whose rows are
// sorted left-to-right and where each row begins after the previous row
// ends. One binary search over the virtual flattened array suffices.
func SearchMatrix(m [][]int, target int) bool {
	if len(m) == 0 || len(m[0]) == 0 {
		return false
	}
	rows, cols := len(m), len(m[0])
	lo, hi := 0, rows*cols-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		v := m[mid/cols][mid%cols]
		switch {
		case v == target:
			return true
		case v < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return false
}

func main() {
	m := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(SearchMatrix(m, 3), SearchMatrix(m, 13))
}
