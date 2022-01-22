package solution1332

import "testing"

func TestRemovePalindromeSub(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 1", "ababa", 1},
		{"test 2", "abb", 2},
		{"test 3", "baabb", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removePalindromeSub(tt.s); got != tt.want {
				t.Errorf("removePalindromeSub() = %d, want %d", got, tt.want)
			}
		})
	}
}
