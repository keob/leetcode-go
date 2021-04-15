package solution198

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{1, 2, 3, 1}, 4},
		{"test 2", []int{2, 7, 9, 3, 1}, 12},
		{"test 3", []int{0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob() = %d, want %d", got, tt.want)
			}
		})
	}
}
