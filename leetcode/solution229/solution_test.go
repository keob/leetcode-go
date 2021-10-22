package solution229

import (
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test 1", []int{3, 2, 3}, []int{3}},
		{"test 2", []int{1}, []int{1}},
		{"test 3", []int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
