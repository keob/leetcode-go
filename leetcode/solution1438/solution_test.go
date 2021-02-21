package solution1438

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		limit int
		want  int
	}{
		{"case 1", []int{8, 2, 4, 7}, 4, 2},
		{"case 2", []int{10, 1, 2, 4, 7, 2}, 5, 4},
		{"case 3", []int{4, 2, 2, 2, 4, 4, 2, 2}, 0, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.nums, tt.limit); got != tt.want {
				t.Errorf("longestSubarray() = %d, want %d", got, tt.want)
			}
		})
	}
}
