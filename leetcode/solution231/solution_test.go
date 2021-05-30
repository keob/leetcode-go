package solution231

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"test 1", 1, true},
		{"test 2", 16, true},
		{"test 3", 3, false},
		{"test 4", 4, true},
		{"test 5", 5, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %t, want %t", got, tt.want)
			}
		})
	}
}
