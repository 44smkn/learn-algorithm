package datastructure

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLinkedList(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		instruction []string
		want        []int
	}{
		{
			"case1",
			[]string{"insert 5", "insert 2", "insert 3", "insert 1", "delete 3", "insert 6", "delete 5"},
			[]int{6, 1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := printList(tt.instruction)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
