package solution1588

import "testing"

func TestSumOddLengthSubarrays(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"test 1", []int{1, 4, 2, 5, 3}, 58},
		{"test 2", []int{1, 2}, 3},
		{"test 3", []int{10, 11, 12}, 66},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOddLengthSubarrays(tt.arr); got != tt.want {
				t.Errorf("sumOddLengthSubarrays() = %d, want %d", got, tt.want)
			}
		})
	}
}
