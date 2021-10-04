package solution482

import "testing"

func TestLicenseKeyFormatting(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{"test 1", "5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
		{"test 2", "2-5g-3-J", 2, "2-5G-3J"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.s, tt.k); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %s, want %s", got, tt.want)
			}
		})
	}
}
