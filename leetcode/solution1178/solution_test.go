package solution1178

import (
	"reflect"
	"testing"
)

func TestFindNumOfValidWords(t *testing.T) {
	tests := []struct {
		name    string
		words   []string
		puzzles []string
		want    []int
	}{
		{
			"case 1",
			[]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
			[]string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
			[]int{1, 1, 3, 2, 4, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findNumOfValidWords(tt.words, tt.puzzles)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
