package solution1629

import "testing"

func TestSlowestKey(t *testing.T) {
	tests := []struct {
		name         string
		releaseTimes []int
		keysPressed  string
		want         byte
	}{
		{"test 1", []int{9, 29, 49, 50}, "cbcd", 'c'},
		{"test 2", []int{12, 23, 36, 46, 62}, "spuda", 'a'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slowestKey(tt.releaseTimes, tt.keysPressed); got != tt.want {
				t.Errorf("slowestKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
