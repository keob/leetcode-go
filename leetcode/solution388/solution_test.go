package solution388

import "testing"

func TestLengthLongestPath(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			"test 1",
			"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			20,
		},
		{
			"test 2",
			"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthLongestPath(tt.input); got != tt.want {
				t.Errorf("lengthLongestPath() = %d, want %d", got, tt.want)
			}
		})
	}
}
