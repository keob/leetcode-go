package solution313

import "testing"

func TestNthSuperUglyNumber(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		primes []int
		want   int
	}{
		{"test 1", 12, []int{2, 7, 13, 19}, 32},
		{"test 2", 1, []int{2, 3, 5}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := nthSuperUglyNumber(tt.n, tt.primes); got != tt.want {
				t.Errorf("nthSuperUglyNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}
