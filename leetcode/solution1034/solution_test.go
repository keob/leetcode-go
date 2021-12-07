package solution1034

import (
	"reflect"
	"testing"
)

func TestColorBorder(t *testing.T) {
	tests := []struct {
		name  string
		grid  [][]int
		row   int
		col   int
		color int
		want  [][]int
	}{
		{"test 1", [][]int{{1, 1}, {1, 2}}, 0, 0, 3, [][]int{{3, 3}, {3, 2}}},
		{"test 2", [][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3, [][]int{{1, 3, 3}, {2, 3, 3}}},
		{"test 3", [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := colorBorder(tt.grid, tt.row, tt.col, tt.color)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("colorBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
