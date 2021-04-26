package solution1011

import "testing"

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		D       int
		want    int
	}{
		{"test 1", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{"test 2", []int{3, 2, 2, 4, 1, 4}, 3, 6},
		{"test 3", []int{1, 2, 3, 1, 1}, 4, 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.weights, tt.D); got != tt.want {
				t.Errorf("shipWithinDays() = %d, want %d", got, tt.want)
			}
		})
	}
}
