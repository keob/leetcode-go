package solution664

import "testing"

func TestStrangePrinter(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 1", "aaabbb", 2},
		{"test 2", "aba", 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := strangePrinter(tt.s); got != tt.want {
				t.Errorf("strangePrinter() = %d, want %d", got, tt.want)
			}
		})
	}
}
