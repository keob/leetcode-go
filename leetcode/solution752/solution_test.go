package solution752

import "testing"

func TestOpenLock(t *testing.T) {
	tests := []struct {
		name     string
		deadends []string
		target   string
		want     int
	}{
		{"test 1", []string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
		{"test 2", []string{"8888"}, "0009", 1},
		{"test 3", []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
		{"test 4", []string{"0000"}, "8888", -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.deadends, tt.target); got != tt.want {
				t.Errorf("openLock() = %d, want %d", got, tt.want)
			}
		})
	}
}
