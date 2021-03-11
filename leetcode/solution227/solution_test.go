package solution227

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"case 1", "3+2*2", 7},
		{"case 2", " 3/2 ", 1},
		{"case 3", "3+5  / 2 ", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.s); got != tt.want {
				t.Errorf("calculate() = %d, want %d", got, tt.want)
			}
		})
	}
}
