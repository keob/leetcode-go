package solution1036

import "testing"

func TestIsEscapePossible(t *testing.T) {
	tests := []struct {
		name    string
		blocked [][]int
		source  []int
		target  []int
		want    bool
	}{
		{
			"test 1",
			[][]int{{0, 1}, {1, 0}},
			[]int{0, 0},
			[]int{0, 2},
			false,
		},
		{
			"test 2",
			[][]int{{}},
			[]int{0, 0},
			[]int{999999, 999999},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEscapePossible(tt.blocked, tt.source, tt.target); got != tt.want {
				t.Errorf("isEscapePossible() = %t, want %t", got, tt.want)
			}
		})
	}
}
