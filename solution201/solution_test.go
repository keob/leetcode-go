package solution201

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{"case 1", 5, 7, 4},
		{"case 2", 0, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.m, tt.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v want %v", got, tt.want)
			}
		})
	}
}
