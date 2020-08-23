package solution263

import "testing"

func TestIsUgly(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"case 1", 6, true},
		{"case 2", 8, true},
		{"case 3", 14, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.num); got != tt.want {
				t.Errorf("isUgly() = %v want %v", got, tt.want)
			}
		})
	}
}
