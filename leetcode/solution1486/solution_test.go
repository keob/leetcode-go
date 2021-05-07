package solution1486

import "testing"

func TestXorOperation(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		start int
		want  int
	}{
		{"test 1", 5, 0, 8},
		{"test 2", 4, 3, 8},
		{"test 3", 1, 7, 7},
		{"test 4", 10, 5, 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperation(tt.n, tt.start); got != tt.want {
				t.Errorf("xorOperation() = %d, want %d", got, tt.want)
			}
		})
	}
}
