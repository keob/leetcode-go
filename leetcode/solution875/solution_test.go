package solution875

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{"test 1", []int{3, 6, 7, 11}, 8, 4},
		{"test 2", []int{30, 11, 23, 4, 20}, 5, 30},
		{"test 3", []int{30, 11, 23, 4, 20}, 6, 23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.piles, tt.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %d, want %d", got, tt.want)
			}
		})
	}
}
