package solution441

import "testing"

func TestArrangeCoins(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 5, 2},
		{"test 2", 8, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.n); got != tt.want {
				t.Errorf("arrangeCoins() = %d, want %d", got, tt.want)
			}
		})
	}
}
