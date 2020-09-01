package solution273

import "testing"

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"case 1", 123, "One Hundred Twenty Three"},
		{"case 2", 12345, "Twelve Thousand Three Hundred Forty Five"},
		{"case 3", 1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.num); got != tt.want {
				t.Errorf("numberToWords() = %v want %v", got, tt.want)
			}
		})
	}
}
