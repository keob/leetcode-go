package solution526

import "testing"

func TestCountArrangement(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 1, 1},
		{"test 2", 2, 2},
		{"test 3", 5, 10},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := countArrangement(tt.n); got != tt.want {
				t.Errorf("countArrangement() = %d, want %d", got, tt.want)
			}
		})
	}
}
