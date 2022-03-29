package solution2024

import "testing"

func TestMaxConsecutiveAnswers(t *testing.T) {
	tests := []struct {
		name      string
		answerKey string
		k         int
		want      int
	}{
		{"test 1", "TTFF", 2, 4},
		{"test 2", "TFFT", 1, 3},
		{"test 3", "TTFTTFTT", 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.answerKey, tt.k); got != tt.want {
				t.Errorf("maxConsecutiveAnswers() = %d, want %d", got, tt.want)
			}
		})
	}
}
