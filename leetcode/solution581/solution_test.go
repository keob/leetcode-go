package solution581

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{2, 6, 4, 8, 10, 9, 15}, 5},
		{"test 2", []int{1, 2, 3, 4}, 0},
		{"test 3", []int{1}, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %d, want %d", got, tt.want)
			}
		})
	}
}
