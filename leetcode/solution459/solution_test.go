package solution459

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"case 1", "abab", true},
		{"case 2", "aba", false},
		{"case 3", "abcabcabcabc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v want %v", got, tt.want)
			}
		})
	}
}
