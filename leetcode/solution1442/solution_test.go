package solution1442

import "testing"

func TestCountTriplets(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"test 1", []int{2, 3, 1, 6, 7}, 4},
		{"test 2", []int{1, 1, 1, 1, 1}, 10},
		{"test 3", []int{2, 3}, 0},
		{"test 4", []int{1, 3, 5, 7, 9}, 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriplets(tt.arr); got != tt.want {
				t.Errorf("countTriplets() = %d, want %d", got, tt.want)
			}
		})
	}
}
