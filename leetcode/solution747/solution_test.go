package solution747

import "testing"

func TestDominantIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{3, 6, 1, 0}, 1},
		{"test 2", []int{1, 2, 3, 4}, -1},
		{"test 3", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.nums); got != tt.want {
				t.Errorf("dominantIndex() = %d, want %d", got, tt.want)
			}
		})
	}
}
