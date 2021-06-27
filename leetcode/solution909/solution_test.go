package solution909

import "testing"

func TestSnakesAndLadders(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  int
	}{
		{
			"test 1",
			[][]int{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 35, -1, -1, 13, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 15, -1, -1, -1, -1},
			},
			4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := snakesAndLadders(tt.board); got != tt.want {
				t.Errorf("snakesAndLadders() = %d, want %d", got, tt.want)
			}
		})
	}
}
