package solution386

import (
	"reflect"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"test 1", 13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"test 2", 2, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lexicalOrder(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder() = %d, want %d", got, tt.want)
			}
		})
	}
}
