package solution881

import "testing"

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		name   string
		people []int
		limit  int
		want   int
	}{
		{"test 1", []int{1, 2}, 3, 1},
		{"test 2", []int{3, 2, 2, 1}, 3, 3},
		{"test 3", []int{3, 5, 3, 4}, 5, 4},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.people, tt.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %d, want %d", got, tt.want)
			}
		})
	}
}
