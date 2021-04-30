package solution137

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{2, 2, 3, 2}, 3},
		{"test 2", []int{0, 1, 0, 1, 0, 1, 99}, 99},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}
