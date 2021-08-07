package solution457

import "testing"

func TestCircularArrayLoop(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"test 1", []int{2, -1, 1, 2, 2}, true},
		{"test 2", []int{-1, 2}, false},
		{"test 3", []int{-2, 1, -1, -2, -2}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.nums); got != tt.want {
				t.Errorf("circularArrayLoop() = %t, want %t", got, tt.want)
			}
		})
	}
}
