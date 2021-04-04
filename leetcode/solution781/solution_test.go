package solution781

import "testing"

func TestNumRabbits(t *testing.T) {
	tests := []struct {
		name    string
		answers []int
		want    int
	}{
		{"test 1", []int{1, 1, 2}, 5},
		{"test 2", []int{10, 10, 10}, 11},
		{"test 3", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.answers); got != tt.want {
				t.Errorf("numRabbits () = %d, want %d", got, tt.want)
			}
		})
	}
}
