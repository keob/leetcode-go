package solution66

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"test 1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"test 2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"test 3", []int{0}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
