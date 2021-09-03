package utils

import (
	"reflect"
	"testing"
)

func TestFlat(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			"test 1",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"test 2",
			[][]int{{}},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Flat(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flat() = %v, want %v", got, tt.want)
			}
		})
	}
}
