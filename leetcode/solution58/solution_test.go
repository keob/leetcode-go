package solution58

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 1", "Hello World", 5},
		{"test 2", "   fly me   to   the moon  ", 4},
		{"test 3", "luffy is still joyboy", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %d, want %d", got, tt.want)
			}
		})
	}
}
