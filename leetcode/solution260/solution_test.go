package solution260

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test 1", []int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{"test 2", []int{-1, 0}, []int{-1, 0}},
		{"test 3", []int{0, 1}, []int{1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
