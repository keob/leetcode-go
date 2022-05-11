package solution942

import (
	"reflect"
	"testing"
)

func TestDiStringMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"test 1", "IDID", []int{0, 4, 1, 3, 2}},
		{"test 2", "III", []int{0, 1, 2, 3}},
		{"test 3", "DDI", []int{3, 2, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diStringMatch(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
