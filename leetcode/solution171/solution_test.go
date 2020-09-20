package solution171

import "testing"

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"cases 1", "A", 1},
		{"cases 2", "AB", 28},
		{"cases 3", "ZY", 701},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.s); got != tt.want {
				t.Errorf("titleToNumber() = %v want %v", got, tt.want)
			}
		})
	}
}
