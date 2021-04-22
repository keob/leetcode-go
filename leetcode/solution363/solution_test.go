package solution363

import "testing"

func TestMaxSumSubmatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		k      int
		want   int
	}{
		{"test 1", [][]int{{1, 0, 1}, {0, -2, 3}}, 2, 2},
		{"test 2", [][]int{{2, 2, -1}}, 3, 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumSubmatrix(tt.matrix, tt.k); got != tt.want {
				t.Errorf("maxSumSubmatrix() = %d, want %d", got, tt.want)
			}
		})
	}
}
