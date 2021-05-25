package solution1787

import "testing"

func TestMinChanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"test 1", []int{1, 2, 0, 3, 0}, 1, 3},
		{"test 2", []int{3, 4, 5, 2, 1, 7, 3, 4}, 3, 3},
		{"test 3", []int{1, 2, 4, 1, 2, 5, 1, 2, 6}, 3, 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.nums, tt.k); got != tt.want {
				t.Errorf("minChanges() = %d, want %d", got, tt.want)
			}
		})
	}
}
