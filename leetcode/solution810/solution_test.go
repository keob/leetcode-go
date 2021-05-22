package solution810

import "testing"

func TestXorGame(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"test 1", []int{1, 1, 2}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := xorGame(tt.nums); got != tt.want {
				t.Errorf("xorgame() = %t, want %t", got, tt.want)
			}
		})
	}
}
