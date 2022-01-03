package solution1185

import "testing"

func TestDayOfTheWeek(t *testing.T) {
	tests := []struct {
		name  string
		day   int
		month int
		year  int
		want  string
	}{
		{"test 1", 31, 8, 2019, "Saturday"},
		{"test 2", 18, 7, 1999, "Sunday"},
		{"test 3", 15, 8, 1993, "Sunday"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfTheWeek(tt.day, tt.month, tt.year); got != tt.want {
				t.Errorf("dayOfTheWeek() = %s, want %s", got, tt.want)
			}
		})
	}
}
