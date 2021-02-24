package solution832

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		name string
		A    [][]int
		want [][]int
	}{
		{
			"case 1",
			[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			"case 2",
			[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flipAndInvertImage(tt.A)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want = %v", got, tt.want)
			}
		})
	}
}
