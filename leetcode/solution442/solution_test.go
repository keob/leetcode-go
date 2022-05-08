package solution442

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test 1", []int{4, 3, 2, 7, 8, 2, 3, 1}, []int{3, 2}},
		{"test 2", []int{1, 1, 2}, []int{1}},
		{"test 3", []int{1}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicates(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
