package solution204

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"case 1", 10, 4},
		{"case 2", 0, 0},
		{"case 3", 100, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.n); got != tt.want {
				t.Errorf("countPrimes() = %v want %v", got, tt.want)
			}
		})
	}
}
