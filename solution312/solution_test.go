package solution312

import "testing"

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"case 1", []int{3, 1, 5, 8}, 167},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
