package offer14I

import "testing"

func TestCuttingRope(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 2, 1},
		{"test 2", 10, 36},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope(tt.n); got != tt.want {
				t.Errorf("cuttingRope() = %d, want %d", got, tt.want)
			}
		})
	}
}
