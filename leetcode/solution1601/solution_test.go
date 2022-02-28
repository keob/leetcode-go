package solution1601

import "testing"

func TestMaximumRequests(t *testing.T) {
	tests := []struct {
		name string
		n    int
		nums [][]int
		want int
	}{
		{"test 1", 5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}, 5},
		{"test 2", 3, [][]int{{0, 0}, {1, 2}, {2, 1}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRequests(tt.n, tt.nums); got != tt.want {
				t.Errorf("maximumRequests() = %d, want %d", got, tt.want)
			}
		})
	}
}
