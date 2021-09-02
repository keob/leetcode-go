package offer20

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"test 1", "0", true},
		{"test 2", "e", false},
		{"test 3", ".", false},
		{"test 4", "    .1  ", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.s); got != tt.want {
				t.Errorf("isNumber() = %t, want %t", got, tt.want)
			}
		})
	}
}
