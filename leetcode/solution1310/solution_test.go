package solution1310

import (
	"reflect"
	"testing"
)

func TestXorQueries(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		queries [][]int
		want    []int
	}{
		{
			"test 1",
			[]int{1, 3, 4, 8},
			[][]int{
				{0, 1},
				{1, 2},
				{0, 3},
				{3, 3},
			},
			[]int{2, 7, 14, 8},
		},
		{
			"test 2",
			[]int{4, 8, 2, 10},
			[][]int{
				{2, 3},
				{1, 3},
				{0, 0},
				{0, 3},
			},
			[]int{8, 0, 4, 4},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := xorQueries(tt.arr, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xorQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
