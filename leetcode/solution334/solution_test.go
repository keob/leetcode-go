package solution334

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"test 1", []int{1, 2, 3, 4, 5}, true},
		{"test 2", []int{5, 4, 3, 2, 1}, false},
		{"test 3", []int{2, 1, 5, 0, 4, 6}, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %t want %t", got, tt.want)
			}
		})
	}
}
