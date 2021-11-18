package solution397

import "testing"

func TestIntegerReplacement(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 8, 3},
		{"test 2", 7, 4},
		{"test 3", 4, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.n); got != tt.want {
				t.Errorf("integerReplacement() = %d, want %d", got, tt.want)
			}
		})
	}
}
