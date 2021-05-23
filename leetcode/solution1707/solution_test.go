package solution1707

import (
	"reflect"
	"testing"
)

func TestMaximizeXor(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    []int
	}{
		{
			"test 1",
			[]int{0, 1, 2, 3, 4},
			[][]int{{3, 1}, {1, 3}, {5, 6}},
			[]int{3, 3, 7},
		},
		{
			"test 2",
			[]int{5, 2, 4, 6, 6, 3},
			[][]int{{12, 4}, {8, 1}, {6, 3}},
			[]int{15, -1, 5},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := maximizeXor(tt.nums, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximizeXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
