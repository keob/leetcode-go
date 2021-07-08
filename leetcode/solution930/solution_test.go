package solution930

import "testing"

func TestNumSubarraysWithSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		goal int
		want int
	}{
		{"test 1", []int{1, 0, 1, 0, 1}, 2, 4},
		{"test 2", []int{0, 0, 0, 0, 0}, 0, 15},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := numSubarraysWithSum(tt.nums, tt.goal)
			if got != tt.want {
				t.Errorf("numSubarraysWithSum() = %d, want %d", got, tt.want)
			}
		})
	}
}
