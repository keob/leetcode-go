package solution395

import "testing"

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"case 1", "aaabb", 3, 3},
		{"case 2", "ababbc", 2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.s, tt.k); got != tt.want {
				t.Errorf("longestSubstring() = %d, want %d", got, tt.want)
			}
		})
	}
}
