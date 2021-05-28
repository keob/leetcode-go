package solution477

import "testing"

func TestTotalHammingDistance(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{4, 14, 2}, 6},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := totalHammingDistance(tt.nums); got != tt.want {
				t.Errorf("totalHammingDistance() = %d, want %d", got, tt.want)
			}
		})
	}
}
