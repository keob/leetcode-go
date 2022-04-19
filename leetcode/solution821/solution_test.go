package solution821

import (
	"reflect"
	"testing"
)

func TestShortestToChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		c    byte
		want []int
	}{
		{"test 1", "loveleetcode", 'e', []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{"test 2", "aaab", 'b', []int{3, 2, 1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestToChar(tt.s, tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
