package solution421

import "testing"

func TestFindMaximumXOR(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{3, 10, 5, 25, 2, 8}, 28},
		{"test 2", []int{0}, 0},
		{"test 3", []int{2, 4}, 6},
		{"test 4", []int{8, 10, 2}, 10},
		{"test 5", []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}, 127},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumXOR(tt.nums); got != tt.want {
				t.Errorf("findMaximumXOR() = %d, want %d", got, tt.want)
			}
		})
	}
}
