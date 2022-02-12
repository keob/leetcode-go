package solution1020

import "testing"

func TestNumEnclaves(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			"test 1",
			[][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			3,
		},
		{
			"test 2",
			[][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEnclaves(tt.grid); got != tt.want {
				t.Errorf("numEnclaves() = %d, want %d", got, tt.want)
			}
		})
	}
}
