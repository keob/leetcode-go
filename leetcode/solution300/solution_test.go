package solution300

import "testing"

func TestLengthOfLTS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"test 2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"test 3", []int{7, 7, 7, 7, 7, 7, 7}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %d, want %d", got, tt.want)
			}
		})
	}
}
