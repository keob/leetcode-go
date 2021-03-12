package solution331

import "testing"

func TestIsValidSerialization(t *testing.T) {
	tests := []struct {
		name     string
		preorder string
		want     bool
	}{
		{"case 1", "9,3,4,#,#,1,#,#,2,#,6,#,#", true},
		{"case 2", "1,#", false},
		{"case 3", "9,#,#,1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSerialization(tt.preorder); got != tt.want {
				t.Errorf("isValidSerialization() = %t, want %t", got, tt.want)
			}
		})
	}
}
