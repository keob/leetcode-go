package s169

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"case 1", []int{3, 2, 3}, 3},
		{"case 2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
