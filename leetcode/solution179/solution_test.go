package solution179

import "testing"

func TestLargestNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{"test 1", []int{10, 2}, "210"},
		{"test 2", []int{3, 30, 34, 5, 9}, "9534330"},
		{"test 3", []int{1}, "1"},
		{"test 4", []int{10}, "10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.nums); got != tt.want {
				t.Errorf("largestNumber() = %s, want %s", got, tt.want)
			}
		})
	}
}
