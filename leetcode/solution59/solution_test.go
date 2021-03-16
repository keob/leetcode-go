package solution59

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			"case 1",
			3,
			[][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			"case 2",
			1,
			[][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateMatrix(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, %v", got, tt.want)
			}
		})
	}
}
