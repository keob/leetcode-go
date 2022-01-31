package solution1342

import "testing"

func TestNumberOfSteps(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"test 1", 14, 6},
		{"test 2", 8, 4},
		{"test 3", 123, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.num); got != tt.want {
				t.Errorf("numberOfSteps() = %d, want %d", got, tt.want)
			}
		})
	}
}
