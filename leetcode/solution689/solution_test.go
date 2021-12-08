package solution689

import (
	"reflect"
	"testing"
)

func TestMaxSumOfThreeSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"test 1", []int{1, 2, 1, 2, 6, 7, 5, 1}, 2, []int{0, 3, 5}},
		{"test 2", []int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2, []int{0, 2, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSumOfThreeSubarrays(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSumOfThreeSubarrays = %v, want %v", got, tt.want)
			}
		})
	}
}
