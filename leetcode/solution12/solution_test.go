package solution12

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"test 1", 3, "III"},
		{"test 2", 4, "IV"},
		{"test 3", 9, "IX"},
		{"test 4", 58, "LVIII"},
		{"test 5", 1994, "MCMXCIV"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.num); got != tt.want {
				t.Errorf("intToRoman() = %s, want %s", got, tt.want)
			}
		})
	}
}
