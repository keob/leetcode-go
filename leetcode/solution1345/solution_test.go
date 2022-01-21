package solution1345

import "testing"

func TestMinJumps(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"test 1", []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, 3},
		{"test 2", []int{7}, 0},
		{"test 3", []int{7, 6, 9, 6, 9, 6, 9, 7}, 1},
		{"test 4", []int{6, 1, 9}, 2},
		{"test 5", []int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minJumps(tt.arr); got != tt.want {
				t.Errorf("minJumps() = %d, want %d", got, tt.want)
			}
		})
	}
}
