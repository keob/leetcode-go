package solution1688

import "testing"

func TestNumberOfMatches(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 7, 6},
		{"test 2", 14, 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfMatches(tt.n); got != tt.want {
				t.Errorf("numberOfMatches() = %d, want %d", got, tt.want)
			}
		})
	}
}
