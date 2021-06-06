package solution474

import "testing"

func TestFindMaxForm(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		m    int
		n    int
		want int
	}{
		{"test 1", []string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{"test 2", []string{"10", "0", "1"}, 1, 1, 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.strs, tt.m, tt.n); got != tt.want {
				t.Errorf("findMaxForm() = %d, want %d", got, tt.want)
			}
		})
	}
}
