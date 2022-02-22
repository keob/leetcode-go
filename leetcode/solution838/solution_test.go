package solution838

import "testing"

func TestPushDominoes(t *testing.T) {
	tests := []struct {
		name     string
		dominoes string
		want     string
	}{
		{"test 1", "RR.L", "RR.L"},
		{"test 2", ".L.R...LR..L..", "LL.RR.LLRRLL.."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %s, want %s", got, tt.want)
			}
		})
	}
}
