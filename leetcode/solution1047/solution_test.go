package solution1047

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"case 1", "abbaca", "ca"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.S); got != tt.want {
				t.Errorf("removeDuplicates() = %s, want %s", got, tt.want)
			}
		})
	}
}
