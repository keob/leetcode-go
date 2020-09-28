package solution342

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"case 1", 16, true},
		{"case 2", 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.num); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
