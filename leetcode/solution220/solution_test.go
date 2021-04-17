package solution220

import "testing"

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		t    int
		want bool
	}{
		{"test 1", []int{1, 2, 3, 1}, 3, 0, true},
		{"test 2", []int{1, 0, 1, 1}, 1, 2, true},
		{"test 3", []int{1, 5, 9, 1, 5, 9}, 2, 3, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %t, want %t", got, tt.want)
			}
		})
	}
}
