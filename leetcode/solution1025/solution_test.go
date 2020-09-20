package solution1025

import "testing"

func TestDivisorGame(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want bool
	}{
		{"case 1", 2, true},
		{"case 2", 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame(tt.N); got != tt.want {
				t.Errorf("divisorGame() = %v want %v", got, tt.want)
			}
		})
	}
}
