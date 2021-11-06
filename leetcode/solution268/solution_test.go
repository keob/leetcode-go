package solution268

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{3, 0, 1}, 2},
		{"test 2", []int{0, 1}, 2},
		{"test 3", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{"test 4", []int{0}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.nums); got != tt.want {
				t.Errorf("missingNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}
