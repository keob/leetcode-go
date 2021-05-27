package solution461

import "testing"

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"test 1", 1, 4, 2},
		{"test 2", 2, 3, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.x, tt.y); got != tt.want {
				t.Errorf("hammingDistance() = %d, want %d", got, tt.want)
			}
		})
	}
}
