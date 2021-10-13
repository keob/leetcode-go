package solution412

import (
	"strconv"
	"strings"
)

func fizzBuzz(n int) (res []string) {
	for i := 1; i <= n; i++ {
		sb := &strings.Builder{}
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		res = append(res, sb.String())
	}
	return
}
