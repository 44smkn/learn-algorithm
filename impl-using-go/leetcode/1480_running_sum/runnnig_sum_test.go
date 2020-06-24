package running_sum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRunningSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}}}

	for _, tt := range tests {
		got := runningSum(tt.input)
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Errorf(diff)
		}
	}
}
