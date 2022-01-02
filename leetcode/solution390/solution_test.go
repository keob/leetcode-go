package solution390

import "testing"

func TestLastRemaining(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 9, 6},
		{"test 2", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastRemaining(tt.n); got != tt.want {
				t.Errorf("lastRemaining() = %d, want %d", got, tt.want)
			}
		})
	}
}
