package offer10I

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 2, 1},
		{"test 2", 5, 5},
		{"test 3", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.n); got != tt.want {
				t.Errorf("fib() = %d, want %d", got, tt.want)
			}
		})
	}
}
