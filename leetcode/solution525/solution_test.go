package solution525

import "testing"

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{0, 1}, 2},
		{"test 2", []int{0, 1, 0}, 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.nums); got != tt.want {
				t.Errorf("findMaxLength() = %d, want %d", got, tt.want)
			}
		})
	}
}
