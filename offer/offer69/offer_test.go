package offer69

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"test 1", []int{0, 1, 0}, 1},
		{"test 2", []int{1, 3, 5, 4, 2}, 2},
		{"test 3", []int{0, 10, 5, 2}, 1},
		{"test 4", []int{3, 4, 5, 1}, 2},
		{"test 5", []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := peakIndexInMountainArray(tt.arr)
			if got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %d, want %d", got, tt.want)
			}
		})
	}
}
