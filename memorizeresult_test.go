package main

import (
	mem "memorizeresult"
	"testing"
)

type testcase struct {
	got    int
	result int
}

func TestCachemem(t *testing.T) {
	cases := []testcase{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{50, 12586269025},
		{55, 139583862445},
	}
	var fibMem mem.Memorized
	fibMem = mem.Cachemem(func(in int) int {
		switch in {
		case 0:
			return 0
		case 1, 2:
			return 1
		default:
			return fibMem(in-2) + fibMem(in-1)
		}
	})
	for _, c := range cases {
		result := fibMem(c.got)

		if result != c.result {
			t.Fatalf("want %v got %v", c.result, result)
		}
	}
}
