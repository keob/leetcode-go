package solution539

import "testing"

func TestFindMinDifference(t *testing.T) {
	tests := []struct {
		name       string
		timePoints []string
		want       int
	}{
		{"test 1", []string{"23:59", "00:00"}, 1},
		{"test 2", []string{"00:00", "23:59", "00:00"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinDifference(tt.timePoints); got != tt.want {
				t.Errorf("findMinDifference() = %d, want %d", got, tt.want)
			}
		})
	}
}
