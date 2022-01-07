package solution1614

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 1", "(1+(2*3)+((8)/4))+1", 3},
		{"test 2", "(1)+((2))+(((3)))", 3},
		{"test 3", "1+(2*3)/(2-1)", 1},
		{"test 4", "1", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.s); got != tt.want {
				t.Errorf("maxDepth() = %d, want %d", got, tt.want)
			}
		})
	}
}
