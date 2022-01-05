package solution1576

import "testing"

func TestModifyString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test 1", "?zs", "azs"},
		{"test 2", "ubv?w", "ubvaw"},
		{"test 3", "j?qg??b", "jaqgacb"},
		{"test 4", "??yw?ipkj?", "abywaipkja"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifyString(tt.s); got != tt.want {
				t.Errorf("modifyString() = %s, want %s", got, tt.want)
			}
		})
	}
}
