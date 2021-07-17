package offer42

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"test 2", []int{1}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %d, want %d", got, tt.want)
			}
		})
	}
}
