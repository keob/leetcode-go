package solution436

import (
	"reflect"
	"testing"
)

func TestFindRightInterval(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      []int
	}{
		{"test 1", [][]int{{1, 2}}, []int{-1}},
		{"test 2", [][]int{{3, 4}, {2, 3}, {1, 2}}, []int{-1, 0, 1}},
		{"test 3", [][]int{{1, 4}, {2, 3}, {3, 4}}, []int{-1, 2, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRightInterval(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRightInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
