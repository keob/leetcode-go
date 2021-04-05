package solution372

import "testing"

func TestSuperPow(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    []int
		want int
	}{
		{"test 1", 2, []int{3}, 8},
		{"test 2", 2, []int{1, 0}, 1024},
		{"test 3", 1, []int{4, 3, 3, 8, 5, 2}, 1},
		{"test 4", 2147483647, []int{2, 0, 0}, 1198},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPow(tt.a, tt.b); got != tt.want {
				t.Errorf("superPow() = %d, want %d", got, tt.want)
			}
		})
	}
}
