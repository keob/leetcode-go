package solution633

import "testing"

func TestJudgeSquareSum(t *testing.T) {
	tests := []struct {
		name string
		c    int
		want bool
	}{
		{"test 1", 5, true},
		{"test 2", 3, false},
		{"test 3", 4, true},
		{"test 4", 2, true},
		{"test 5", 1, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %t, want %t", got, tt.want)
			}
		})
	}
}
