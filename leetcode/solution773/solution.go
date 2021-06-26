package solution773

import "strings"

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

type pair struct {
	status string
	step   int
}

func slidingPuzzle(board [][]int) int {
	const target = "123450"
	s := make([]byte, 0, 6)

	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}

	start := string(s)
	if start == target {
		return 0
	}

	get := func(status string) (ret []string) {
		s := []byte(status)
		x := strings.Index(status, "0")
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			ret = append(ret, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}

	q := []pair{{start, 0}}
	seen := map[string]bool{start: true}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}

	return -1
}
