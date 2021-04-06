package solution80

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{1, 1, 1, 2, 2, 3}, 5},
		{"test 2", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates = %d, want %d", got, tt.want)
			}
		})
	}
}
