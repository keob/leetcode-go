package solution453

import "testing"

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{1, 2, 3}, 3},
		{"test 2", []int{1, 1, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.nums); got != tt.want {
				t.Errorf("minMoves() = %d, want %d", got, tt.want)
			}
		})
	}
}
