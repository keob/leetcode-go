package solution913

import "testing"

func TestCatMouseGame(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  int
	}{
		{
			"test 1",
			[][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}},
			0,
		},
		{
			"test 2",
			[][]int{{1, 3}, {0}, {3}, {0, 2}},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := catMouseGame(tt.graph); got != tt.want {
				t.Errorf("catMouseGame() = %d, want %d", got, tt.want)
			}
		})
	}
}
