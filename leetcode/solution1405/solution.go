package solution1405

import "sort"

func longestDiverseString(a int, b int, c int) string {
	res := []byte{}
	cnt := []struct {
		c  int
		ch byte
	}{{a, 'a'}, {b, 'b'}, {c, 'c'}}

	for {
		sort.Slice(cnt, func(i, j int) bool { return cnt[i].c > cnt[j].c })
		hasNext := false
		for i, p := range cnt {
			if p.c <= 0 {
				break
			}
			m := len(res)
			if m >= 2 && res[m-2] == p.ch && res[m-1] == p.ch {
				continue
			}
			hasNext = true
			res = append(res, p.ch)
			cnt[i].c--
			break
		}
		if !hasNext {
			return string(res)
		}
	}
}
