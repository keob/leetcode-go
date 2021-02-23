package solution1052

import "testing"

func TestMaxSatisfied(t *testing.T) {
	tests := []struct {
		name      string
		customers []int
		grumpy    []int
		X         int
		want      int
	}{
		{"case 1", []int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied(tt.customers, tt.grumpy, tt.X); got != tt.want {
				t.Errorf("maxSatisfied() = %d, want %d", got, tt.want)
			}
		})
	}
}
