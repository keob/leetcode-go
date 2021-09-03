package solution52

import "testing"

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 4, 2},
		{"test 2", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.n); got != tt.want {
				t.Errorf("totalNQueens() = %d, want %d", got, tt.want)
			}
		})
	}
}
