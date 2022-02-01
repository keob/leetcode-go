package solution1763

import "testing"

func TestLongestNiceSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test 1", "YazaAay", "aAa"},
		{"test 2", "Bb", "Bb"},
		{"test 3", "c", ""},
		{"test 4", "dDzeE", "dD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubstring(tt.s); got != tt.want {
				t.Errorf("longestNiceSubstring() = %s, got %s", got, tt.want)
			}
		})
	}
}
