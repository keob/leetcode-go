package solution1706

import (
	"reflect"
	"testing"
)

func TestFindBall(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want []int
	}{
		{
			"test 1",
			[][]int{
				{1, 1, 1, -1, -1},
				{1, 1, 1, -1, -1},
				{-1, -1, -1, 1, 1},
				{1, 1, 1, 1, -1},
				{-1, -1, -1, -1, -1},
			},
			[]int{1, -1, -1, -1, -1},
		},
		{
			"test 2",
			[][]int{{-1}},
			[]int{-1},
		},
		{
			"test 3",
			[][]int{
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
			},
			[]int{0, 1, 2, 3, 4, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findBall(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
