package solution492

import (
	"reflect"
	"testing"
)

func TestConstructRectangle(t *testing.T) {
	tests := []struct {
		name string
		area int
		want []int
	}{
		{"test 1", 4, []int{2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructRectangle(tt.area)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
