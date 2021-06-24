package solution149

import "testing"

func TestMaxPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{"test 1", [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
		{"test 2", [][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.points); got != tt.want {
				t.Errorf("maxPoints() = %d, want %d", got, tt.want)
			}
		})
	}
}
