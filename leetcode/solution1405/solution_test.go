package solution1405

import "testing"

func TestLongestDiverseString(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want string
	}{
		{"test 1", 1, 1, 7, "ccaccbcc"},
		{"test 2", 2, 2, 1, "abbac"},
		{"test 3", 7, 1, 0, "aabaa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDiverseString(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("longestDiverseString() = %s, want %s", got, tt.want)
			}
		})
	}
}
