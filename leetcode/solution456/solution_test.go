package solution456

import "testing"

func TestFind132pattern(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"case 1", []int{1, 2, 3, 4}, false},
		{"case 2", []int{3, 1, 4, 2}, true},
		{"case 3", []int{-1, 3, 2, 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.nums); got != tt.want {
				t.Errorf("find132pattern() = %t, want %t", got, tt.want)
			}
		})
	}
}
