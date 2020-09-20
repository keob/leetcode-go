package solution290

import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		str     string
		want    bool
	}{
		{"case 1", "abba", "dog cat cat dog", true},
		{"case 2", "abba", "dog cat cat fish", false},
		{"case 3", "aaaa", "dog cat cat dog", false},
		{"case 4", "abba", "dog dog dog dog", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.pattern, tt.str); got != tt.want {
				t.Errorf("wordPattern() = %v want %v", got, tt.want)
			}
		})
	}
}
