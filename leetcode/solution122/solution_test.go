package solution122

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"name 1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"name 2", []int{1, 2, 3, 4, 5}, 4},
		{"name 3", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v want %v", got, tt.want)
			}
		})
	}
}
