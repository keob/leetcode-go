package solution81

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"case 1", []int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{"case 2", []int{2, 5, 6, 0, 0, 1, 2}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %t, want %t", got, tt.want)
			}
		})
	}
}
