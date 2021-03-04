package solution354

import "testing"

func TestMaxEnvelopes(t *testing.T) {
	tests := []struct {
		name      string
		envelopes [][]int
		want      int
	}{
		{"case 1", [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{"case 2", [][]int{{1, 1}, {1, 1}, {1, 1}}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %d, want %d", got, tt.want)
			}
		})
	}
}
