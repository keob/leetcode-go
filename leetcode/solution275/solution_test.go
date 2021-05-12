package solution275

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{"test 1", []int{0, 1, 3, 5, 6}, 3},
		{"test 2", []int{}, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.citations); got != tt.want {
				t.Errorf("hIndex() = %d, want %d", got, tt.want)
			}
		})
	}
}
