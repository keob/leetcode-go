package solution713

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"test 1", []int{10, 5, 2, 6}, 100, 8},
		{"test 2", []int{1, 2, 3}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.nums, tt.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %d, want %d", got, tt.want)
			}
		})
	}
}
