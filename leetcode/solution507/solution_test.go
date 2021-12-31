package solution507

import "testing"

func TestCheckPerfectNumber(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"test 1", 28, true},
		{"test 2", 6, true},
		{"test 3", 496, true},
		{"test 4", 8128, true},
		{"test 3", 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %t, want %t", got, tt.want)
			}
		})
	}
}
