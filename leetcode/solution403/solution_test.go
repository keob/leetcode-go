package solution403

import "testing"

func TestCanCross(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   bool
	}{
		{"test 1", []int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{"test 2", []int{0, 1, 2, 3, 4, 8, 9, 11}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := canCross(tt.stones); got != tt.want {
				t.Errorf("cancross() = %t, want %t", got, tt.want)
			}
		})
	}
}
