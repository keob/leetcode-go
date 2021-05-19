package solution1738

import "testing"

func TestKthLargestValue(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		k      int
		want   int
	}{
		{"test 1", [][]int{{5, 2}, {1, 6}}, 1, 7},
		{"test 2", [][]int{{5, 2}, {1, 6}}, 2, 5},
		{"test 3", [][]int{{5, 2}, {1, 6}}, 3, 4},
		{"test 4", [][]int{{5, 2}, {1, 6}}, 4, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestValue(tt.matrix, tt.k); got != tt.want {
				t.Errorf("kthLargestValue() = %d, want %d", got, tt.want)
			}
		})
	}
}
