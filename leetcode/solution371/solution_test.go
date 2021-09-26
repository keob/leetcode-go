package solution371

import "testing"

func TestGetSum(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"test 1", 1, 2, 3},
		{"test 2", 2, 3, 5},
		{"test 3", -1, 1, 0},
		{"test 4", -1, -1, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.a, tt.b); got != tt.want {
				t.Errorf("getSum() = %d, want %d", got, tt.want)
			}
		})
	}
}
