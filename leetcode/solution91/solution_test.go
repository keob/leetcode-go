package solution91

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 1", "12", 2},
		{"test 2", "226", 3},
		{"test 3", "0", 0},
		{"test 4", "06", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings() = %d, want %d", got, tt.want)
			}
		})
	}
}
