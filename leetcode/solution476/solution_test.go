package solution476

import "testing"

func TestFindComplement(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"test 1", 5, 2},
		{"test 2", 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findComplement(tt.num); got != tt.want {
				t.Errorf("findComplement() = %d, want %d", got, tt.want)
			}
		})
	}
}
