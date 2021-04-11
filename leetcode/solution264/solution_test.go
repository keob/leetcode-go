package solution264

import "testing"

func TestNthUglyNumbers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 10, 12},
		{"test 2", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}
