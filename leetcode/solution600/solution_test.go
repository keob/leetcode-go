package solution600

import "testing"

func TestFindIntegers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 1, 2},
		{"test 2", 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntegers(tt.n); got != tt.want {
				t.Errorf("findIntegers() = %d, want %d", got, tt.want)
			}
		})
	}
}
