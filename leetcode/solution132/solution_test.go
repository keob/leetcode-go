package solution132

import "testing"

func TestMinCut(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"case 1", "aab", 1},
		{"case 2", "a", 0},
		{"case 3", "ab", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.s); got != tt.want {
				t.Errorf("minCut() = %d, want %d", got, tt.want)
			}
		})
	}
}
