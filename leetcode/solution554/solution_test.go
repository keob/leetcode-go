package solution554

import "testing"

func TestLeastBricks(t *testing.T) {
	tests := []struct {
		name string
		wall [][]int
		want int
	}{
		{
			"test 1",
			[][]int{
				{1, 2, 2, 1},
				{3, 1, 2},
				{1, 3, 2},
				{2, 4},
				{3, 1, 2},
				{1, 3, 1, 1},
			},
			2,
		},
		{
			"test 2",
			[][]int{
				{1},
				{1},
				{1},
			},
			3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := leastBricks(tt.wall); got != tt.want {
				t.Errorf("leastBricks(0 = %d, want %d", got, tt.want)
			}
		})
	}
}
