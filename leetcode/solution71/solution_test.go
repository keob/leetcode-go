package solution71

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"test 1", "/home/", "/home"},
		{"test 2", "/../", "/"},
		{"test 3", "/home//foo", "/home/foo"},
		{"test 4", "/a/./b/../../c/", "/c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %s, want %s", got, tt.want)
			}
		})
	}
}
