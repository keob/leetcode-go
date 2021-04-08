package solution153

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{3, 4, 5, 1, 2}, 1},
		{"test 2", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"test 3", []int{11, 13, 15, 17}, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin = %d, want %d", got, tt.want)
			}
		})
	}
}
