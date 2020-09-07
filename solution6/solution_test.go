package solution6

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{"name 1", "LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN"},
		{"name 2", "LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.s, tt.numRows); got != tt.want {
				t.Errorf("convert () = %v, want %v", got, tt.want)
			}
		})
	}
}
