package solution282

import (
	"reflect"
	"testing"
)

func TestAddOperators(t *testing.T) {
	tests := []struct {
		name   string
		num    string
		target int
		want   []string
	}{
		{"test 1", "123", 6, []string{"1+2+3", "1*2*3"}},
		{"test 2", "232", 8, []string{"2+3*2", "2*3+2"}},
		{"test 3", "105", 5, []string{"1*0+5", "10-5"}},
		{"test 4", "00", 0, []string{"0+0", "0-0", "0*0"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addOperators(tt.num, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, got %v", got, tt.want)
			}
		})
	}
}
