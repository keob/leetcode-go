package solution31

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test 1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"test 2", []int{3, 2, 1}, []int{1, 2, 3}},
		{"test 3", []int{1, 1, 5}, []int{1, 5, 1}},
		{"test 4", []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("nextPermutation() = %v, got %v", tt.nums, tt.want)
			}
		})
	}
}
