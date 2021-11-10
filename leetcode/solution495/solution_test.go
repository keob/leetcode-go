package solution495

import "testing"

func TestFindPoisonedDuration(t *testing.T) {
	tests := []struct {
		name       string
		timeseries []int
		duration   int
		want       int
	}{
		{"test 1", []int{1, 4}, 2, 4},
		{"test 2", []int{1, 2}, 2, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPoisonedDuration(tt.timeseries, tt.duration)
			if got != tt.want {
				t.Errorf("findPoisonedDuration() = %d, want %d", got, tt.want)
			}
		})
	}
}
