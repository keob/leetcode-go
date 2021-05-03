package solution7

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"test 1", 123, 321},
		{"test 2", -123, -321},
		{"test 3", 120, 21},
		{"test 4", 0, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %d, want %d", got, tt.want)
			}
		})
	}
}
