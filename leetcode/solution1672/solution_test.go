package solution1672

import "testing"

func TestMaximumWealth(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]int
		want     int
	}{
		{"test 1", [][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{"test 2", [][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
		{"test 3", [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWealth(tt.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %d, want %d", got, tt.want)
			}
		})
	}
}
