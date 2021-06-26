package solution773

import "testing"

func TestSlidingPuzzle(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  int
	}{
		{"test 1", [][]int{{1, 2, 3}, {4, 0, 5}}, 1},
		{"test 2", [][]int{{1, 2, 3}, {5, 4, 0}}, -1},
		{"test 3", [][]int{{4, 1, 2}, {5, 0, 3}}, 5},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingPuzzle(tt.board); got != tt.want {
				t.Errorf("slidingPuzzle() = %d, want %d", got, tt.want)
			}
		})
	}
}
