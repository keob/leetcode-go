package solution890

import (
	"reflect"
	"testing"
)

func TestFindAndReplacePattern(t *testing.T) {
	tests := []struct {
		name    string
		words   []string
		pattern string
		want    []string
	}{
		{
			"test 1",
			[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
			"abb",
			[]string{"mee", "aqq"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAndReplacePattern(tt.words, tt.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
