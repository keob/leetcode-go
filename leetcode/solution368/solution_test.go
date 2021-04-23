package solution368

import (
	"reflect"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test 1", []int{1, 2, 3}, []int{2, 1}},
		{"test 2", []int{1, 2, 4, 8}, []int{8, 4, 2, 1}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := largestDivisibleSubset(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
