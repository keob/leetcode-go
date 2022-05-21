package solution961

import "testing"

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{1, 2, 3, 3}, 3},
		{"test 2", []int{2, 1, 2, 5, 3, 2}, 2},
		{"test 3", []int{5, 1, 5, 2, 5, 3, 5, 4}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedNTimes(tt.nums); got != tt.want {
				t.Errorf("repeatedNTimes() = %d, want %d", got, tt.want)
			}
		})
	}
}
