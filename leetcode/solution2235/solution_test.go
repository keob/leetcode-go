package solution2235

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		want int
	}{
		{"test 1", 12, 5, 17},
		{"test 2", -10, 4, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.num1, tt.num2); got != tt.want {
				t.Errorf("sum() = %d, want %d", got, tt.want)
			}
		})
	}
}
