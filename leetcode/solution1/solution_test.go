package solution1

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"test 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"test 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"test 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
