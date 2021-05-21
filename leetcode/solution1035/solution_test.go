package solution1035

import "testing"

func TestMaxUncrossedLines(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{"test 1", []int{1, 4, 2}, []int{1, 2, 4}, 2},
		{"test 2", []int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
		{"test 3", []int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("maxUncrossedLines() = %d, want %d", got, tt.want)
			}
		})
	}
}
