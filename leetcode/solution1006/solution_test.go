package solution1006

import "testing"

func TestClumsy(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"case 1", 4, 7},
		{"case 2", 10, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsy(tt.N); got != tt.want {
				t.Errorf("clumsy() = %d, want %d", got, tt.want)
			}
		})
	}
}
