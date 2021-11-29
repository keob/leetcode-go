package solution786

import (
	"reflect"
	"testing"
)

func TestKthSmallestPrimeFraction(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want []int
	}{
		{"test 1", []int{1, 2, 3, 5}, 3, []int{2, 5}},
		{"test 2", []int{1, 7}, 1, []int{1, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthSmallestPrimeFraction(tt.arr, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
