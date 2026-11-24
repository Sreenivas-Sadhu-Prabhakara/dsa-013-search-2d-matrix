package main

import "testing"

func TestSearchMatrix(t *testing.T) {
	m := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	cases := []struct {
		target int
		want   bool
	}{
		{3, true}, {13, false}, {1, true}, {60, true}, {0, false}, {61, false}, {11, true},
	}
	for _, c := range cases {
		if got := SearchMatrix(m, c.target); got != c.want {
			t.Errorf("SearchMatrix(m, %d) = %v, want %v", c.target, got, c.want)
		}
	}
	if SearchMatrix([][]int{}, 1) {
		t.Error("empty matrix should be false")
	}
	if !SearchMatrix([][]int{{5}}, 5) {
		t.Error("single cell hit should be true")
	}
}
