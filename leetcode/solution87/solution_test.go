package solution87

import "testing"

func TestIsScramble(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{"test 1", "great", "rgeat", true},
		{"test 2", "abcde", "caebd", false},
		{"test 3", "a", "a", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.s1, tt.s2); got != tt.want {
				t.Errorf("isScramble() = %t, want %t", got, tt.want)
			}
		})
	}
}
