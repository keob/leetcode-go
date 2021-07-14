package solution1818

import "testing"

func TestMinAbsoluteSumDiff(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{"test 1", []int{1, 7, 5}, []int{2, 3, 5}, 3},
		{"test 2", []int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}, 0},
		{"test 3", []int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}, 20},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := minAbsoluteSumDiff(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("minAbsoluteSumDiff() = %d, want %d", got, tt.want)
			}
		})
	}
}
