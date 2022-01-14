package solution373

import (
	"reflect"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		{
			"test 1",
			[]int{1, 7, 11},
			[]int{2, 4, 6},
			3,
			[][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			"test 2",
			[]int{1, 1, 2},
			[]int{1, 2, 3},
			2,
			[][]int{{1, 1}, {1, 1}},
		},
		{
			"test 3",
			[]int{1, 2},
			[]int{3},
			3,
			[][]int{{1, 3}, {2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kSmallestPairs(tt.nums1, tt.nums2, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
