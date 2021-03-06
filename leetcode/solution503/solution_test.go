package solution503

import (
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"case 1", []int{1, 2, 1}, []int{2, -1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextGreaterElements(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
