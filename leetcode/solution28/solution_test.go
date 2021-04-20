package solution28

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{"test 1", "hello", "ll", 2},
		{"test 2", "aaaaa", "bba", -1},
		{"test 3", "", "", 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %d, want %d", got, tt.want)
			}
		})
	}
}
