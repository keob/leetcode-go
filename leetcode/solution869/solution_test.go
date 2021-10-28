package solution869

import "testing"

func TestReorderedPowerOf2(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"test 1", 1, true},
		{"test 2", 10, false},
		{"test 3", 16, true},
		{"test 4", 24, false},
		{"test 5", 46, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.n); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %t, want %t", got, tt.want)
			}
		})
	}
}
