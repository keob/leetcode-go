package solution115

import "testing"

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{"case 1", "rabbbit", "rabbit", 3},
		{"case 2", "babgbag", "bag", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.s, tt.t); got != tt.want {
				t.Errorf("numDistinct() = %d, want %d", got, tt.want)
			}
		})
	}
}
