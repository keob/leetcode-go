package solution1436

import "testing"

func TestDestCity(t *testing.T) {
	tests := []struct {
		name  string
		paths [][]string
		want  string
	}{
		{"test 1", [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}, "Sao Paulo"},
		{"test 2", [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}, "A"},
		{"test 3", [][]string{{"A", "Z"}}, "Z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCity(tt.paths); got != tt.want {
				t.Errorf("destCity() = %s, want %s", got, tt.want)
			}
		})
	}
}
