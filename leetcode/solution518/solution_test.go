package solution518

import "testing"

func TestChange(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{
		{"test 1", 5, []int{1, 2, 5}, 4},
		{"test 2", 3, []int{2}, 0},
		{"test 3", 10, []int{10}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := change(tt.amount, tt.coins)
			if got != tt.want {
				t.Errorf("change() = %d, want %d", got, tt.want)
			}
		})
	}
}
