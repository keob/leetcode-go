package solution451

import "testing"

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test 1", "tree", "eetr"},
		{"test 2", "cccaaa", "cccaaa"},
		{"test 3", "Aabb", "bbAa"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.s); got != tt.want {
				t.Errorf("frequencySort() = %s, want %s", got, tt.want)
			}
		})
	}
}
