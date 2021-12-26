package solution1078

import (
	"reflect"
	"testing"
)

func TestFindOcurrences(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		first  string
		second string
		want   []string
	}{
		{
			"test 1",
			"alice is a good girl she is a good student",
			"a",
			"good",
			[]string{"girl", "student"},
		},
		{
			"test 2",
			"we will we will rock you",
			"we",
			"will",
			[]string{"we", "rock"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOcurrences(tt.text, tt.first, tt.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOcurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
