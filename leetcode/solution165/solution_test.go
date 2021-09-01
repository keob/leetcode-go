package solution165

import "testing"

func TestCompareVersion(t *testing.T) {
	tests := []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{
		{"test 1", "1.01", "1.001", 0},
		{"test 2", "1.0", "1.0.0", 0},
		{"test 3", "0.1", "1.1", -1},
		{"test 4", "1.0.1", "1", 1},
		{"test 5", "7.5.2.4", "7.5.3", -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.version1, tt.version2); got != tt.want {
				t.Errorf("compareVersion() = %d, want %d", got, tt.want)
			}
		})
	}
}
