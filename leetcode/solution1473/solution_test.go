package solution1473

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		name   string
		houses []int
		cost   [][]int
		m      int
		n      int
		target int
		want   int
	}{
		{
			"test 1", []int{0, 0, 0, 0, 0},
			[][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}},
			5, 2, 3, 9,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.houses, tt.cost, tt.m, tt.n, tt.target); got != tt.want {
				t.Errorf("minCost() = %d, want %d", got, tt.want)
			}
		})
	}
}
