package solution233

import "testing"

func TestCountDigitOne(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 13, 6},
		{"test 2", 0, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigitOne(tt.n); got != tt.want {
				t.Errorf("countDigitOne() = %d, want %d", got, tt.want)
			}
		})
	}
}
