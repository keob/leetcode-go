package solution1137

import "testing"

func TestTribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 4, 4},
		{"test 2", 1, 1},
		{"test 3", 25, 1389537},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.n); got != tt.want {
				t.Errorf("tribonacci() = %d, want %d", got, tt.want)
			}
		})
	}
}
