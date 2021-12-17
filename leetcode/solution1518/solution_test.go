package solution1518

import "testing"

func TestNumWaterBottles(t *testing.T) {
	tests := []struct {
		name        string
		numBottles  int
		numExchange int
		want        int
	}{
		{"test 1", 9, 3, 13},
		{"test 2", 15, 4, 19},
		{"test 3", 5, 5, 6},
		{"test 4", 2, 3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWaterBottles(tt.numBottles, tt.numExchange); got != tt.want {
				t.Errorf("numWaterBottles() = %d, want %d", got, tt.want)
			}
		})
	}
}
