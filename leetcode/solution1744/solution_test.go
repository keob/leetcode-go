package solution1744

import (
	"reflect"
	"testing"
)

func TestCanEat(t *testing.T) {
	tests := []struct {
		name         string
		candiesCount []int
		queries      [][]int
		want         []bool
	}{
		{
			"test 1",
			[]int{7, 4, 5, 3, 8},
			[][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 1000000000}},
			[]bool{true, false, true},
		},
		{
			"test 2",
			[]int{5, 2, 6, 4, 1},
			[][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}},
			[]bool{false, true, true, false, false},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := canEat(tt.candiesCount, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canEat() = %v, want %v", got, tt.want)
			}
		})
	}
}
