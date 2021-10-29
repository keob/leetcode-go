package solution335

import "testing"

func TestIsSelfCrossing(t *testing.T) {
	tests := []struct {
		name     string
		distance []int
		want     bool
	}{
		{"test 1", []int{1, 1, 1, 1}, true},
		{"test 2", []int{2, 1, 1, 2}, true},
		{"test 3", []int{1, 2, 3, 4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSelfCrossing(tt.distance); got != tt.want {
				t.Errorf("isSelfCrossing() = %t, want %t", got, tt.want)
			}
		})
	}
}
