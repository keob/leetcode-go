package solution172

import "testing"

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 3, 0},
		{"test 2", 5, 1},
		{"test 3", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.n); got != tt.want {
				t.Errorf("trailingZeroes() = %d, want %d", got, tt.want)
			}
		})
	}
}
