package solution375

import "testing"

func TestGetMoneyAmount(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 10, 16},
		{"test 2", 1, 0},
		{"test 3", 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneyAmount(tt.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %d, want %d", got, tt.want)
			}
		})
	}
}
