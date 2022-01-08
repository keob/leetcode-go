package solution89

import (
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"test 1", 2, []int{0, 1, 3, 2}},
		{"test 2", 1, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grayCode(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
