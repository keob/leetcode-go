package solution136

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{2, 2, 1}, 1},
		{"test 2", []int{4, 1, 2, 1, 2}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}
