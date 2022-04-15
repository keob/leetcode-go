package solution357

import "testing"

func TestCountNumbersWithUniqueDigits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 2, 91},
		{"test 2", 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumbersWithUniqueDigits(tt.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits() = %d, want %d", got, tt.want)
			}
		})
	}
}
