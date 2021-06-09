package solution879

import "testing"

func TestProfitableSchemes(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		minProfit int
		group     []int
		profit    []int
		want      int
	}{
		{"test 1", 5, 3, []int{2, 2}, []int{2, 3}, 2},
		{"test 2", 10, 5, []int{2, 3, 5}, []int{6, 7, 8}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := profitableSchemes(tt.n, tt.minProfit, tt.group, tt.profit)
			if got != tt.want {
				t.Errorf("profitableSchemes() =  %d, want %d", got, tt.want)
			}
		})
	}
}
