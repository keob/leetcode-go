package solution686

import "testing"

func TestRepeatedStringMatch(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{"test 1", "abcd", "cdabcdab", 3},
		{"test 2", "a", "aa", 2},
		{"test 3", "a", "a", 1},
		{"test 4", "abc", "wxyz", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.a, tt.b); got != tt.want {
				t.Errorf("repeatedStringMatch() = %d, want %d", got, tt.want)
			}
		})
	}
}
