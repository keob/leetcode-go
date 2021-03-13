package solution350

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"case 1", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"case 2", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
