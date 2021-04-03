package solution1143

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{"test 1", "abcde", "ace", 3},
		{"test 2", "abc", "abc", 3},
		{"test 3", "abc", "def", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %d, want %d", got, tt.want)
			}
		})
	}
}
