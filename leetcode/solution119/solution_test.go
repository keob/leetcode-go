package solution119

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{"case 1", 3, []int{1, 3, 3, 1}},
		{"case 2", 4, []int{1, 4, 6, 4, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRow(tt.rowIndex)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
