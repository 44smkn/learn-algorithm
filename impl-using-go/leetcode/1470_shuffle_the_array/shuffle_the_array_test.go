package shuffle_the_array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		// input
		nums []int
		n    int
		// output
		want []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, tt := range tests {
		got := shuffle(tt.nums, tt.n)
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Error(diff)
		}
	}
}
