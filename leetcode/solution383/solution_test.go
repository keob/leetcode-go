package solution383

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{"case 1", "a", "b", false},
		{"case 2", "aa", "ab", false},
		{"case 3", "aa", "aab", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("canConstruct() = %t want %t", got, tt.want)
			}
		})
	}
}
