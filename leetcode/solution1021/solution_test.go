package solution1021

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test 1", "(()())(())", "()()()"},
		{"test 2", "(()())(())(()(()))", "()()()()(())"},
		{"test 3", "()()", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.s); got != tt.want {
				t.Errorf("removeOuterParentheses() = %s, want %s", got, tt.want)
			}
		})
	}
}
