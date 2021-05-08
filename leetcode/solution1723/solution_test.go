package solution1723

import "testing"

func TestMinimumTimeRequired(t *testing.T) {
	tests := []struct {
		name string
		jobs []int
		k    int
		want int
	}{
		{"test 1", []int{3, 2, 3}, 3, 3},
		{"test 2", []int{1, 2, 4, 7, 8}, 2, 11},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeRequired(tt.jobs, tt.k); got != tt.want {
				t.Errorf("minimumTimeRequired() = %d, want %d", got, tt.want)
			}
		})
	}
}
