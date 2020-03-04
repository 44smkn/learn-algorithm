package main

import (
	"fmt"
	"testing"
)

func TestSortNumAscending(t *testing.T) {
	arr := []int{5, 1, 4, 3, 2, 7, 9, 8, 6, 10}
	gs := []int{4, 1}

	ShellSort(arr, gs)
	if fmt.Sprint(arr) != "[1 2 3 4 5 6 7 8 9 10]" {
		fmt.Println(arr)
		t.Fatal("not expected array")
	}
}
