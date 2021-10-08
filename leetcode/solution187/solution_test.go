package solution187

import (
	"reflect"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"test 1", "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"test 2", "AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRepeatedDnaSequences(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
