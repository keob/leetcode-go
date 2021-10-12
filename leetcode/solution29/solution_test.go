package solution29

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{"test 1", 10, 3, 3},
		{"test 2", 7, -3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.dividend, tt.divisor); got != tt.want {
				t.Errorf("divide() = %d, want %d", got, tt.want)
			}
		})
	}
}
