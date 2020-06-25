package elementaryalignment_test

import (
	"aoj/elementaryalignment"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShellSort(t *testing.T) {
	arr := []int{5, 1, 4, 3, 2, 7, 9, 8, 6, 10}
	gs := []int{4, 1}

	elementaryalignment.ShellSort(arr, gs)

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if diff := cmp.Diff(want, arr); diff != "" {
		t.Error(diff)
	}
}
