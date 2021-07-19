package solution1838

import "testing"

func TestMaxFrequency(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"test 1", []int{1, 2, 4}, 5, 3},
		{"test 2", []int{1, 4, 8, 13}, 5, 2},
		{"test 3", []int{3, 9, 6}, 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxFrequency() = %d want %d", got, tt.want)
			}
		})
	}
}
