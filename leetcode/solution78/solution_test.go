package solution78

import (
	"reflect"
	"testing"

	"github.com/keob/leetcode/utils"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			"test 1",
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}},
		},
		{
			"test 2",
			[]int{0},
			[][]int{{0}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			if !reflect.DeepEqual(utils.Flat(got), utils.Flat(tt.want)) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
