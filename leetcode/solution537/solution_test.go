package solution537

import "testing"

func TestComplexNumberMultiply(t *testing.T) {
	tests := []struct {
		name string
		num1 string
		num2 string
		want string
	}{
		{"test 1", "1+1i", "1+1i", "0+2i"},
		{"test 2", "1+-1i", "1+-1i", "0+-2i"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.num1, tt.num2); got != tt.want {
				t.Errorf("complexNumberMultiply() = %s, want %s", got, tt.want)
			}
		})
	}
}
