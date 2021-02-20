package solution697

import "testing"

func TestFindShortestSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"case 1", []int{1, 2, 2, 3, 1}, 2},
		{"case 2", []int{1, 2, 2, 3, 1, 4, 2}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
