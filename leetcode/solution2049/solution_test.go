package solution2049

import "testing"

func TestCountHighestScoreNodes(t *testing.T) {
	tests := []struct {
		name    string
		parents []int
		want    int
	}{
		{"test 1", []int{-1, 2, 0, 2, 0}, 3},
		{"test 2", []int{-1, 2, 0}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHighestScoreNodes(tt.parents); got != tt.want {
				t.Errorf("countHighestScoreNodes() = %d, want %d", got, tt.want)
			}
		})
	}
}
