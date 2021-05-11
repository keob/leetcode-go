package solution1734

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		encoded []int
		want    []int
	}{
		{"test 1", []int{3, 1}, []int{1, 2, 3}},
		{"test 2", []int{6, 5, 4, 6}, []int{2, 4, 1, 5, 3}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := decode(tt.encoded)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
