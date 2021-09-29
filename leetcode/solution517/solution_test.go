package solution517

import "testing"

func TestFindMinMoves(t *testing.T) {
	tests := []struct {
		name     string
		machines []int
		want     int
	}{
		{"test 1", []int{1, 0, 5}, 3},
		{"test 2", []int{0, 3, 0}, 2},
		{"test 3", []int{0, 2, 0}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinMoves(tt.machines); got != tt.want {
				t.Errorf("findMinMoves() = %d, got %d", got, tt.want)
			}
		})
	}
}
