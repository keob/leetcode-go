package solution523

import "testing"

func TestCheckSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"test 1", []int{23, 2, 4, 6, 7}, 6, true},
		{"test 2", []int{23, 2, 6, 4, 7}, 6, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %t, want %t", got, tt.want)
			}
		})
	}
}
