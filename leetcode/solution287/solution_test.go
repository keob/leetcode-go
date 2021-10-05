package solution287

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{1, 3, 4, 2, 2}, 2},
		{"test 2", []int{3, 1, 3, 4, 2}, 3},
		{"test 3", []int{1, 1}, 1},
		{"test 4", []int{1, 1, 2}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.nums); got != tt.want {
				t.Errorf("findDuplicate() = %d, want %d", got, tt.want)
			}
		})
	}
}
