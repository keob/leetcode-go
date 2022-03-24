package solution440

import "testing"

func TestFindKthNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{"test 1", 13, 2, 10},
		{"test 2", 1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.n, tt.k); got != tt.want {
				t.Errorf("findKthNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}
