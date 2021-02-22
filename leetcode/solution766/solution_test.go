package solution766

import "testing"

func TestIsToeplitzMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   bool
	}{
		{"case 1", [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true},
		{"case 2", [][]int{{1, 2}, {2, 2}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToeplitzMatrix(tt.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrix() = %t\n want = %t", got, tt.want)
			}
		})
	}
}
