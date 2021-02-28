package solution896

import "testing"

func TestIsMonotonic(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want bool
	}{
		{"case 1", []int{1, 2, 2, 3}, true},
		{"case 2", []int{6, 5, 4, 4}, true},
		{"case 3", []int{1, 3, 2}, false},
		{"case 4", []int{1, 2, 4, 5}, true},
		{"case 5", []int{1, 1, 1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.A); got != tt.want {
				t.Errorf("isMonotonic() = %t, want %t", got, tt.want)
			}
		})
	}
}
