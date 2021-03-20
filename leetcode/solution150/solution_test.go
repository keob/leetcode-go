package solution150

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			"case 1",
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
		{
			"case 2",
			[]string{"4", "13", "5", "/", "+"},
			6,
		},
		{
			"case 3",
		[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tokens); got != tt.want {
				t.Errorf("evalRPN() = %d, want %d", got, tt.want)
			}
		})
	}
}
