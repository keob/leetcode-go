package solution2016

import "testing"

func TestMaximumDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{7, 1, 5, 4}, 4},
		{"test 2", []int{9, 4, 3, 2}, -1},
		{"test 3", []int{1, 5, 2, 10}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDifference(tt.nums); got != tt.want {
				t.Errorf("maximumDifference() = %d, want %d", got, tt.want)
			}
		})
	}
}
