package solution639

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 1", "*", 9},
		{"test 2", "1*", 18},
		{"test 3", "2*", 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings() = %d, want %d", got, tt.want)
			}
		})
	}
}
