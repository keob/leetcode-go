package solution524

import "testing"

func TestFindLongestWord(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		dictionary []string
		want       string
	}{
		{"test 1", "abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"test 2", "abpcplea", []string{"a", "b", "c"}, "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.s, tt.dictionary); got != tt.want {
				t.Errorf("findLongestWord() = %s, want %s", got, tt.want)
			}
		})
	}
}
