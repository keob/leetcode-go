package solution292

import "testing"

func TestCanWinNim(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"test 1", 4, false},
		{"test 2", 1, true},
		{"test 3", 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.n); got != tt.want {
				t.Errorf("canWinNim() = %t, want %t", got, tt.want)
			}
		})
	}
}
