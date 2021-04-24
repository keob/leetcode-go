package solution377

import "testing"

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"test 1", []int{1, 2, 3}, 4, 7},
		{"test 2", []int{9}, 3, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum4(tt.nums, tt.target); got != tt.want {
				t.Errorf("combinationSum4() = %d, want %d", got, tt.want)
			}
		})
	}
}
