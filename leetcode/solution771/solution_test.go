package solution771

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		name string
		J    string
		S    string
		want int
	}{
		{"case 1", "aA", "aAAbbbb", 3},
		{"case 2", "z", "ZZ", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.J, tt.S); got != tt.want {
				t.Errorf("numJewelsInStones() = %v want %v", got, tt.want)
			}
		})
	}
}
