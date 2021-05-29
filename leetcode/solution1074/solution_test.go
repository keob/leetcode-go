package solution1074

import "testing"

func TestNumSubmatrixSumTarget(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   int
	}{
		{"test 1", [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0, 4},
		{"test 2", [][]int{{1, -1}, {-1, 1}}, 0, 5},
		{"test 3", [][]int{{904}}, 0, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubmatrixSumTarget(tt.matrix, tt.target); got != tt.want {
				t.Errorf("numSubmatrixSumTarget() = %d, want %d", got, tt.want)
			}
		})
	}
}
