package solution504

import "testing"

func TestConvertToBase7(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"test 1", 100, "202"},
		{"test 2", -7, "-10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.num); got != tt.want {
				t.Errorf("convertToBase7() = %s, want %s", got, tt.want)
			}
		})
	}
}
