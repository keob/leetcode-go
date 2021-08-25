package solution797

import (
	"reflect"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  [][]int
	}{
		{
			"test 1",
			[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			[][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
		{
			"test 2",
			[][]int{{1}, {}},
			[][]int{{0, 1}},
		},
		{
			"test 3",
			[][]int{{1, 2, 3}, {2}, {3}, {}},
			[][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}},
		},
		{
			"test 4",
			[][]int{{1, 3}, {2}, {3}, {}},
			[][]int{{0, 1, 2, 3}, {0, 3}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := allPathsSourceTarget(tt.graph)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
