package solution1725

import "testing"

func TestCountGoodRectangles(t *testing.T) {
	tests := []struct {
		name       string
		rectangles [][]int
		want       int
	}{
		{
			"test 1",
			[][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}},
			3,
		},
		{
			"test 2",
			[][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodRectangles(tt.rectangles); got != tt.want {
				t.Errorf("countGoodRectangles() = %d, want %d", got, tt.want)
			}
		})
	}
}
