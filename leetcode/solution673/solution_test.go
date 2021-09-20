package solution673

import "testing"

func TestFindNumberOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{1, 3, 5, 4, 7}, 2},
		{"test 2", []int{2, 2, 2, 2, 2}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %d, want %d", got, tt.want)
			}
		})
	}
}
