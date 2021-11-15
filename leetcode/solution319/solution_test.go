package solution319

import "testing"

func TestBulbSwitch(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test 1", 3, 1},
		{"test 2", 0, 0},
		{"test 3", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bulbSwitch(tt.n); got != tt.want {
				t.Errorf("bulbSwitch = %d, want %d", got, tt.want)
			}
		})
	}
}
