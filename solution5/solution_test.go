package solution5

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"name 1", "babad", "bab"},
		{"name 2", "cbbd", "bb"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v want %v", got, tt.want)
			}
		})
	}
}
