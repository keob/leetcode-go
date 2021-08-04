package solution611

import "testing"

func TestTriangleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{2, 2, 3, 4}, 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(tt.nums); got != tt.want {
				t.Errorf("triangleNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}
