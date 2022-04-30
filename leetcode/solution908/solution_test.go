package solution908

import "testing"

func TestSmallestRange(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"test 1", []int{1}, 0, 0},
		{"test 2", []int{0, 10}, 2, 6},
		{"test 3", []int{1, 3, 6}, 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRange(tt.nums, tt.k); got != tt.want {
				t.Errorf("smallestRange() = %d, want %d", got, tt.want)
			}
		})
	}
}
