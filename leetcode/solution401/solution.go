package solution401

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) (res []string) {
	for i := 0; i < 1024; i++ {
		h, m := i>>6, i&63
		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn {
			res = append(res, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return
}
