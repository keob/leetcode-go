package solution1720

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		encoded []int
		first   int
		want    []int
	}{
		{
			"test 1",
			[]int{1, 2, 3},
			1,
			[]int{1, 0, 2, 1},
		},
		{
			"test 2",
			[]int{6, 2, 7, 3},
			4,
			[]int{4, 2, 0, 7, 4},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := decode(tt.encoded, tt.first)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
