package solution224

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"case 1", "1 + 1", 2},
		{"case 2", "2 - 1 + 2", 3},
		{"case 3", "(1+(4+5+2)-3)+(6+8)", 23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.s); got != tt.want {
				t.Errorf("calculate() = %d, want %d", got, tt.want)
			}
		})
	}
}
