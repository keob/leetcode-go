package solution458

import "testing"

func TestPoorPigs(t *testing.T) {
	tests := []struct {
		name          string
		buckets       int
		minutesToDie  int
		minutesToTest int
		want          int
	}{
		{"test 1", 1000, 15, 60, 5},
		{"test 2", 4, 15, 15, 2},
		{"test 3", 4, 15, 30, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := poorPigs(tt.buckets, tt.minutesToDie, tt.minutesToTest)
			if got != tt.want {
				t.Errorf("poorPigs() = %d, want %d", got, tt.want)
			}
		})
	}
}
