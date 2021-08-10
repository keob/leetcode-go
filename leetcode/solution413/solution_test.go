package solution413

import "testing"

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{1, 2, 3, 4}, 3},
		{"test 2", []int{1}, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %d, want %d", got, tt.want)
			}
		})
	}
}
