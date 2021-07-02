package solution1833

import "testing"

func TestMaxIceCream(t *testing.T) {
	tests := []struct {
		name  string
		costs []int
		coins int
		want  int
	}{
		{"test 1", []int{1, 3, 2, 4, 1}, 7, 4},
		{"test 2", []int{10, 6, 8, 7, 7, 8}, 5, 0},
		{"test 3", []int{1, 6, 3, 1, 2, 5}, 20, 6},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIceCream(tt.costs, tt.coins); got != tt.want {
				t.Errorf("maxIceCream() = %d, want %d", got, tt.want)
			}
		})
	}
}
