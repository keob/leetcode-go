package solution868

import "testing"

func TestBinaryGap(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 22, 2},
		{"test 2", 8, 0},
		{"test 3", 5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.n); got != tt.want {
				t.Errorf("binaryGap() = %d, want %d", got, tt.want)
			}
		})
	}
}
