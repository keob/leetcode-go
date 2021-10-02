package solution405

import "testing"

func TestToHex(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"test 1", 26, "1a"},
		{"test 2", -1, "ffffffff"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.num); got != tt.want {
				t.Errorf("toHex() = %s, want %s", got, tt.want)
			}
		})
	}
}
