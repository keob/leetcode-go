package solution883

import "testing"

func TestProjectionArea(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"test 1", [][]int{{1, 2}, {3, 4}}, 17},
		{"test 2", [][]int{{1, 0}, {0, 2}}, 8},
		{"test 3", [][]int{{2}}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := projectionArea(tt.grid); got != tt.want {
				t.Errorf("projectionArea() = %d, want %d", got, tt.want)
			}
		})
	}
}
