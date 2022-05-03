package solution905

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test 1", []int{0}, []int{0}},
		{"test 2", []int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArrayByParity(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
