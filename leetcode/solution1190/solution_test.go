package solution1190

import "testing"

func TestReverseParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test 1", "(abcd)", "dcba"},
		{"test 2", "(u(love)i)", "iloveu"},
		{"test 3", "(ed(et(oc))el)", "leetcode"},
		{"test 4", "a(bcdefghijkl(mno)p)q", "apmnolkjihgfedcbq"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.s); got != tt.want {
				t.Errorf("reverseParentheses() = %s, %s", got, tt.want)
			}
		})
	}
}
