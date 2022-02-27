package solution553

import "testing"

func TestOptimalDivision(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{"test 1", []int{1000, 100, 10, 2}, "1000/(100/10/2)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.nums); got != tt.want {
				t.Errorf("optimalDivision() = %s, want %s", got, tt.want)
			}
		})
	}
}
