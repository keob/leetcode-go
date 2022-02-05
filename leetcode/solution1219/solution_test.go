package solution1219

import "testing"

func TestGetMaximumGold(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			"test 1",
			[][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}},
			24,
		},
		{
			"test 1",
			[][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}},
			28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGold(tt.grid); got != tt.want {
				t.Errorf("getMaximumGold() = %d, want %d", got, tt.want)
			}
		})
	}
}
