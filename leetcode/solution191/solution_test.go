package solution191

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want int
	}{
		{"case 1", 00000000000000000000000000001011, 3},
		{"case 2", 00000000000000000000000010000000, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.num); got != tt.want {
				t.Errorf("hammingWeight() = %d, want %d", got, tt.want)
			}
		})
	}
}
