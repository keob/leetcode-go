package solution2000

import "testing"

func TestReversePrefix(t *testing.T) {
	tests := []struct {
		name string
		word string
		ch   byte
		want string
	}{
		{"test 1", "abcdefd", 'd', "dcbaefd"},
		{"test 2", "xyxzxe", 'z', "zxyxxe"},
		{"test 3", "abcd", 'z', "abcd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefix(tt.word, tt.ch); got != tt.want {
				t.Errorf("reversePrefix() = %s, want %s", got, tt.want)
			}
		})
	}
}
