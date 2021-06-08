package solution1049

import "testing"

func TestLastStoneWeightII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{"test 1", []int{2, 7, 4, 1, 8, 1}, 1},
		{"test 2", []int{31, 26, 33, 21, 40}, 5},
		{"test 3", []int{1, 2}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeightII(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeightII() = %d, want %d", got, tt.want)
			}
		})
	}
}
