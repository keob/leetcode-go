package solution279

import "testing"

func TestNumSquares(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"test 1", 12, 3},
		{"test 2", 13, 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.x); got != tt.want {
				t.Errorf("numSquares() = %d, want %d", got, tt.want)
			}
		})
	}
}
