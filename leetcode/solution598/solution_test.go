package solution598

import "testing"

func TestMaxCount(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		ops  [][]int
		want int
	}{
		{"test 1", 3, 3, [][]int{{2, 2}, {3, 3}}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.m, tt.n, tt.ops); got != tt.want {
				t.Errorf("maxCount() = %d, want %d", got, tt.want)
			}
		})
	}
}
