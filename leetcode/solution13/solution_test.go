package solution13

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 1", "III", 3},
		{"test 2", "IV", 4},
		{"test 3", "IX", 9},
		{"test 4", "LVIII", 58},
		{"test 5", "MCMXCIV", 1994},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %d, want %d", got, tt.want)
			}
		})
	}
}
