package solution553

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	n := len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	res := &strings.Builder{}
	res.WriteString(fmt.Sprintf("%d/(%d", nums[0], nums[1]))
	for _, num := range nums[2:] {
		res.WriteByte('/')
		res.WriteString(strconv.Itoa(num))
	}
	res.WriteByte(')')
	return res.String()
}
