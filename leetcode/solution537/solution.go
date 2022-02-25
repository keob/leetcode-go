package solution537

import (
	"fmt"
	"strconv"
	"strings"
)

func parseComplexNumber(num string) (real int, imag int) {
	i := strings.IndexByte(num, '+')
	real, _ = strconv.Atoi(num[:i])
	imag, _ = strconv.Atoi(num[i+1 : len(num)-1])
	return
}

func complexNumberMultiply(num1 string, num2 string) string {
	real1, imag1 := parseComplexNumber(num1)
	real2, imag2 := parseComplexNumber(num2)

	return fmt.Sprintf("%d+%di", real1*real2-imag1*imag2, real1*imag2+imag1*real2)
}
