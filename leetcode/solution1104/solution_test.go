package solution1104

import (
	"reflect"
	"testing"
)

func TestPathInZigZagTree(t *testing.T) {
	tests := []struct {
		name  string
		label int
		want  []int
	}{
		{"test 1", 14, []int{1, 3, 4, 14}},
		{"test 2", 26, []int{1, 2, 6, 10, 26}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := pathInZigZagTree(tt.label)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathInZigZagTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
